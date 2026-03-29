package monitor

import (
	"context"
	"sync"
	"time"

	"github.com/CFF4HA/Staircase/internal/engine/core"
	"github.com/CFF4HA/Staircase/internal/engine/database"
	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/google/uuid"
)

// This is the monitor package, which will be responsible for monitoring the state of the staircase and providing
// database updates to the metadata tables.
type Monitor struct {
	wg *sync.WaitGroup
	c  context.Context

	Job chan types.RetrievalJob
}

func Create() *Monitor {

	return &Monitor{
		wg:  &sync.WaitGroup{},
		Job: make(chan types.RetrievalJob, 100),
	}
}

func (m *Monitor) Wait() {
	m.wg.Wait()
}

// This function will enqueue the jobs that are new in the database
// as determined by the `is_initialzied` field.
func qNewJobs() []types.DatabaseJob {
	var jobs []types.DatabaseJob
	var job_ids []uuid.UUID
	db := database.Database()
	if tx := db.Model(&types.DatabaseJobMetadata{}).Where("is_initialized = ?", false).Pluck("job_id", &job_ids); tx.Error != nil {
		core.Logger.Error("Failed to query database for new jobs", "error", tx.Error)
		return nil
	}

	tx := db.Model(&types.DatabaseJob{}).Where("id IN ?", job_ids).Preload("Metadata").Preload("Staircases").Find(&jobs)
	if tx.Error != nil {
		core.Logger.Error("Failed to query database for new jobs", "error", tx.Error)
		return nil
	}
	return jobs
}

// This fuction will enqueue the jobs that are finished in the
// database and have a last scan time which has become stale.
func qRefreshableJobs() []types.DatabaseJob {
	var allJobs []types.DatabaseJob
	var job_ids []uuid.UUID
	db := database.Database()
	if tx := db.Model(&types.DatabaseJobMetadata{}).Where("is_finished = ?", true).Pluck("job_id", &job_ids); tx.Error != nil {
		core.Logger.Error("Failed to query database for finished jobs", "error", tx.Error)
		return nil
	}

	tx := db.Model(&types.DatabaseJob{}).Where("id IN ?", job_ids).Preload("Metadata").Preload("Staircases").Find(&allJobs)
	if tx.Error != nil {
		core.Logger.Error("Failed to query database for finished jobs", "error", tx.Error)
		return nil
	}

	var jobs []types.DatabaseJob
	for _, job := range allJobs {
		scanDeadline := job.Metadata.LastScan.Add(time.Duration(job.Frequency) * time.Hour)
		if scanDeadline.Before(time.Now()) {
			core.Logger.Debug("Job is refreshable, adding to queue", "job_id", job.ID, "scanDeadline", scanDeadline, "now", time.Now())
			jobs = append(jobs, job)
		}
	}

	return jobs
}

func (m *Monitor) worker() {
	defer m.wg.Done()

	ticker := time.NewTicker(time.Duration(core.Config.DatabasePollInterval) * time.Second)
	defer ticker.Stop()

	parser := NewStaircaseProcessor()

	// this function will iterate through the database and
	// collect all new jobs that are available.
	for {
		select {
		case <-ticker.C:
			var jobs []types.DatabaseJob
			jobs = append(jobs, qNewJobs()...)
			jobs = append(jobs, qRefreshableJobs()...)

			for _, job := range jobs {
				// here we want to use the parser to completely parse the
				// jobs and the staircase statements attached to them, converting
				// them into a format that is more convenient for the engine.
				var retrieval types.RetrievalJob
				retrieval.Id = job.ID

				for _, staircase := range job.Staircases {
					q := parser.Process(staircase.Declaration)
					q.Id = staircase.ID

					retrieval.Queries = append(retrieval.Queries, *q)
				}

				{
					db := database.Database()
					job.Metadata.IsInitialized = true
					db.Save(&job.Metadata)
				}

				retrieval.JobMetadata = job.Metadata
				m.Job <- retrieval
			}
		case <-m.c.Done():
			core.Logger.Debug("Worker shutting down...")
			return
		}
	}
}

func (m *Monitor) Start(ctx context.Context, workers int) chan types.RetrievalJob {
	m.c = ctx

	for range workers {
		m.wg.Add(1)
		go m.worker()
	}

	return m.Job
}

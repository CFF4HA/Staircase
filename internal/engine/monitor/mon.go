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

func (m *Monitor) worker() {
	defer m.wg.Done()

	ticker := time.NewTicker(time.Duration(core.Config.DatabasePollInterval) * time.Second)
	defer ticker.Stop()

	db := database.Database()
	parser := NewStaircaseProcessor()

	// this function will iterate through the database and
	// collect all new jobs that are available.
	for {
		select {
		case <-ticker.C:
			var job_ids []uuid.UUID
			if tx := db.Model(&types.DatabaseJobMetadata{}).Where("is_initialized = ?", false).Pluck("job_id", &job_ids); tx.Error != nil {
				core.Logger.Error("Failed to query database for new jobs", "error", tx.Error)
				break
			}

			core.Logger.Debug("Found new jobs", "count", len(job_ids))

			var jobs []types.DatabaseJob
			tx := db.Model(&types.DatabaseJob{}).Where("id IN ?", job_ids).Preload("Metadata").Preload("Staircases").Find(&jobs)
			if tx.Error != nil {
				core.Logger.Error("Failed to query database for new jobs", "error", tx.Error)
			}

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

				job.Metadata.IsInitialized = true
				job.Metadata.LastScan = time.Now()
				db.Save(&job.Metadata)

				core.Logger.Debug("Dispatched new job", "job", retrieval)
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

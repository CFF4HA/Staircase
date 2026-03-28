package chrome

import (
	"context"
	"encoding/base64"
	"strings"
	"sync"

	"github.com/CFF4HA/Staircase/internal/engine/core"
	"github.com/CFF4HA/Staircase/internal/engine/database"
	"github.com/CFF4HA/Staircase/internal/types"

	"github.com/chromedp/chromedp"
)

type WorkerPool struct {
	jobs chan types.RetrievalJob
	wg   *sync.WaitGroup
}

func NewWorkerPool(queue chan types.RetrievalJob) *WorkerPool {

	return &WorkerPool{
		jobs: queue,
		wg:   &sync.WaitGroup{},
	}
}

func (w *WorkerPool) Wait() {
	w.wg.Wait()
}

func (w *WorkerPool) Start(ctx context.Context, workers int) {
	for range workers {
		w.wg.Add(1)
		go w.spawn(ctx)
	}
}

func (w *WorkerPool) spawn(ctx context.Context) {
	defer w.wg.Done()
	c, err := NewClient()
	if err != nil {
		panic(err)
	}

	core.Logger.Info("successfully created new chrome client for worker")

	db := database.Database()
	for {
		select {
		case job := <-w.jobs:
			var results []types.StaircaseEngineLog
			var injections []string

			// at this point we can say that the job is in
			// progress.
			{
				job.JobMetadata.IsInProgress = true
				db.Save(&job.JobMetadata)
			}

			// at this point, we are performing the actual scraping job
			// using the regex engine we built.
			failed := false
			chrome_ctx, _ := chromedp.NewContext(context.Background())
			for _, query := range job.Queries {
				res, err := c.Perform(chrome_ctx, query)
				if err != nil {
					failed = true
					break
				}

				injected, _ := c.QueryToJS(query)
				injections = append(injections, injected)
				results = append(results, *res)
			}

			if failed {
				core.Logger.Error("failed to perform job, marking as failed", "job_id", job.Id)
				continue
			}

			// here we are going to first upload it to the
			// default database (which we need to set-up)
			var screenshot []byte
			if err := chromedp.Run(chrome_ctx, chromedp.FullScreenshot(&screenshot, 90)); err != nil {
				core.Logger.Error("failed to take screenshot of page", "err", err)
				continue
			}

			// at this point we have finished the job, we have the result
			// and we have a screenshot, so there is nothing else to do.
			{
				job.JobMetadata.IsFinished = true
				job.JobMetadata.Image = base64.StdEncoding.EncodeToString(screenshot)
				job.JobMetadata.InjectedJavascript = strings.Join(injections, "\n")
				job.JobMetadata.Result = results[len(results)-1].Result
				db.Save(&job.JobMetadata)
				core.Logger.Info("successfully completed job", "job_id", job.Id)
			}

		case <-ctx.Done():
			core.Logger.Info("worker received shutdown signal, exiting")
		}
	}
}

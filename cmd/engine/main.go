package main

import (
	"context"
	"flag"

	"github.com/CFF4HA/Staircase/internal/engine/chrome"
	"github.com/CFF4HA/Staircase/internal/engine/core"
	"github.com/CFF4HA/Staircase/internal/engine/database"
	"github.com/CFF4HA/Staircase/internal/engine/monitor"
)

func main() {
	db := flag.String("db", "", "the database holding the staircase records")
	poll := flag.Int("poll", 10, "the interval at which to poll the database for new jobs")
	flag.Parse()

	// Initialize the core configuration
	core.Config.DatabasePollInterval = *poll

	// Initialize the database contents
	database.Init(*db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup the engine
	m := monitor.Create()
	jobs := m.Start(ctx, 1)

	// setup the worker pools
	wp := chrome.NewWorkerPool(jobs)
	wp.Start(ctx, 4)

	m.Wait()
	wp.Wait()
}

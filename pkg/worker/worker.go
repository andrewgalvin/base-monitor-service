package worker

import (
	"base-monitor-service/pkg/config"
	"base-monitor-service/pkg/database"
	"base-monitor-service/pkg/httpclient"
	"base-monitor-service/pkg/logger"
	"base-monitor-service/pkg/model"
	"base-monitor-service/pkg/processing"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type WorkerPool struct {
	config        *config.Config
	tasks         chan model.Task // Your "Task" object that each worker will use
	activeWorkers int32
	db            *database.Database
	mutex         sync.Mutex
}

func NewWorkerPool(cfg *config.Config, tasks chan model.Task, db *database.Database) *WorkerPool {
	return &WorkerPool{
		config: cfg,
		tasks:  tasks,
		db:     db,
	}
}

// Start starts the worker pool by creating a ticker that runs every 5 seconds.
// It checks the number of active workers and logs the count.
// If there are no active workers, it shuts down the worker pool by stopping the ticker and closing the tasks channel.
// It also creates worker goroutines based on the configured worker pool size.
func (wp *WorkerPool) Start() {
	ticker := time.NewTicker(time.Second * 5) // X is the number of minutes you want
	go func() {
		for range ticker.C {
			activeWorkers := atomic.LoadInt32(&wp.activeWorkers)
			logger.Info("Active workers:", activeWorkers)

			if activeWorkers == 0 {
				logger.Info("No active workers, shutting down...")
				ticker.Stop()
				close(wp.tasks)
				return
			}
		}
	}()

	for i := 0; i < wp.config.WorkerPoolSize; i++ {
		go wp.worker(i)
	}
}

// worker is a goroutine that processes tasks from the worker pool.
// It selects a random proxy from the pool's configuration and creates an HTTP client with that proxy.
// It then starts the monitor and processes the task using the client.
// If an error occurs during the process, it logs the error message.
func (wp *WorkerPool) worker(id int) {

	logger.Info("Worker [", id, "] started")
	now := time.Now().UnixNano()  // Current time in nanoseconds
	uniqueSeed := now + int64(id) // Combine them to increase uniqueness

	src := rand.NewSource(uniqueSeed)
	rnd := rand.New(src)

	// Loop over the tasks channel and process each task
	for task := range wp.tasks {
		// Use rnd.Intn to select a random proxy index
		proxyIndex := rnd.Intn(len(wp.config.Proxies))
		proxyStr := wp.config.Proxies[proxyIndex]

		client, err := wp.newClientWithProxy(proxyStr)
		if err != nil {
			logger.Error("Failed to create HTTP client for worker ", id, ": ", err)
			return // Consider how you want to handle this error
		}

		// Start the monitor
		logger.Info("Worker [", id, "] processing task: ", task)
		atomic.AddInt32(&wp.activeWorkers, 1)
		defer atomic.AddInt32(&wp.activeWorkers, -1)

		err = processing.ProcessTask(client)
		if err != nil {
			logger.Error("Worker [", id, "] failed to process task")
		}
	}
}

// newClientWithProxy creates a new HTTP client with the specified proxy.
// It takes a proxy string as a parameter and returns a pointer to an http.Client and an error.
// The proxy string should be in the format "host:port".
func (wp *WorkerPool) newClientWithProxy(proxy string) (*http.Client, error) {
	return httpclient.NewClient(wp.config, proxy)
}

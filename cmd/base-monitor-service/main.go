package main

import (
	"base-monitor-service/pkg/config"
	"base-monitor-service/pkg/database"
	"base-monitor-service/pkg/model"
	"base-monitor-service/pkg/worker"
)

func main() {
	// Where your program starts

	// Initialize config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Initialize database
	db, err := database.GetDBInstance(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize task channel
	taskChan := make(chan model.Task, 1000) // 1000 is the buffer size

	// Initialize worker pool
	workerPool := worker.NewWorkerPool(cfg, taskChan, db)
	// Start worker pool
	workerPool.Start()

	// Add logic here to populat the task channel
	// with tasks to be processed by the worker pool
	//
	// Example:
	// task := model.Task{
	// 	// Add task details here
	// }
	// taskChan <- task
	//
	// Ideally, this would be connected to your database
	// and pull data from there to be processed by the worker pool

	// Ensures application never exits
	select {}
}

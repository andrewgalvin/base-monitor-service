package database

import (
	"context"
	"sync"
	"time"

	"base-monitor-service/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

var (
	dbInstance  *Database
	initDBOnce  sync.Once
	initDBError error
)

// initDBClient initializes the database client and connects to the server.
// It takes a pointer to a config.Config struct as a parameter.
// If an error occurs during initialization or connection, it is stored in the initDBError variable.
// The connection is established using the MongoDB URI specified in the config.
// After successful connection, the function creates a new Database instance with the connected client.
func initDBClient(cfg *config.Config) {
	initDBOnce.Do(func() {
		// Create a new client and connect to the server
		client, err := mongo.NewClient(options.Client().ApplyURI(cfg.DBConnectionString))
		if err != nil {
			initDBError = err
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			initDBError = err
			return
		}

		// Optionally, you can check the connection here by pinging the MongoDB server
		err = client.Ping(ctx, nil)
		if err != nil {
			initDBError = err
			return
		}

		dbInstance = &Database{Client: client}
	})
}

// GetDBInstance returns the singleton instance of the Database.
// If the instance is not initialized, it will be initialized using the provided configuration.
// It returns the Database instance and any initialization error encountered.
func GetDBInstance(cfg *config.Config) (*Database, error) {
	if dbInstance == nil {
		initDBClient(cfg)
	}
	return dbInstance, initDBError
}

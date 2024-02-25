package repository

import (
	"base-monitor-service/pkg/database"
)

type StubhubRepository struct {
	db *database.Database
}

func NewStubhubRepository(db *database.Database) *StubhubRepository {
	return &StubhubRepository{db: db}
}

// GetAllItemsFromCollection retrieves all items from a collection in the database.
// It returns an error if the retrieval fails.
// Assumes you are using MongoDB as your database.
func (repo *StubhubRepository) GetAllItemsFromCollection() error {
	// Sample code to retrieve all items from a collection in MongoDB
	// collection := repo.db.Client.Database("your_database").Collection("your_collection")
	// ctx := context.Background()

	// cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }
	// defer cursor.Close(ctx)

	// for cursor.Next(ctx) {
	// 	var event model.YourModel
	// 	if err = cursor.Decode(&event); err != nil {
	// 		log.Fatal(err)
	// 		return nil, err
	// 	}
	// 	events = append(events, event)
	// }

	// if err = cursor.Err(); err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }

	// return events, nil
}

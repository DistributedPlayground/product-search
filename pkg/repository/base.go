package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Base[T any] struct {
	Collection *mongo.Collection
}

// FindOne retrieves a single document based on the provided filter.
func (b Base[T]) FindOne(ctx context.Context, filter bson.M) (*T, error) {
	var result T
	err := b.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

func (b Base[T]) GetById(ctx context.Context, id string) (*T, error) {
	filter := bson.M{"_id": id}
	return b.FindOne(ctx, filter)
}

// List retrieves multiple documents based on the provided filter. You can also provide options for sorting, limiting, etc.
func (b Base[T]) List(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*T, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cur, err := b.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*T
	for cur.Next(ctx) {
		var elem T
		if err := cur.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// Count retrieves the count of documents based on the provided filter.
func (b Base[T]) Count(ctx context.Context, filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	count, err := b.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Aggregate executes an aggregation pipeline and returns the results.
func (b Base[T]) Aggregate(ctx context.Context, pipeline mongo.Pipeline) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cur, err := b.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// Distinct retrieves distinct values for the provided field based on the filter.
func (b Base[T]) Distinct(ctx context.Context, fieldName string, filter bson.M) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	values, err := b.Collection.Distinct(ctx, fieldName, filter)
	if err != nil {
		return nil, err
	}
	return values, nil
}

// Writes

func (b Base[T]) InsertOne(ctx context.Context, document *T) (*mongo.InsertOneResult, error) {
	return b.Collection.InsertOne(ctx, document)
}

func (b Base[T]) InsertMany(ctx context.Context, documents []interface{}) (*mongo.InsertManyResult, error) {
	return b.Collection.InsertMany(ctx, documents)
}

func (b Base[T]) UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return b.Collection.UpdateOne(ctx, filter, update)
}

func (b Base[T]) UpdateMany(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return b.Collection.UpdateMany(ctx, filter, update)
}

func (b Base[T]) DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	return b.Collection.DeleteOne(ctx, filter)
}

func (b Base[T]) DeleteMany(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	return b.Collection.DeleteMany(ctx, filter)
}

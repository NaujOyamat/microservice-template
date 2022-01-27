package mongo

import (
	"context"
	"fmt"
	"log"

	"Github.com/NaujOyamat/microservice-template/internal/domain/courses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseRepository struct {
	db *mongo.Database
}

func NewCourseRepository(db *mongo.Database) *CourseRepository {
	return &CourseRepository{db}
}

func (repo *CourseRepository) Save(ctx context.Context, course *courses.Course) error {
	doc := bson.D{
		primitive.E{Key: "Name", Value: course.Name()},
		primitive.E{Key: "Duration", Value: course.Duration()},
	}

	result, err := repo.db.Collection("Courses").InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	log.Printf("\nDocumentID: %v", result.InsertedID)

	return nil
}

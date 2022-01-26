package postgres

import (
	"context"

	"Github.com/NaujOyamat/microservice-template/internal/domain/courses"
)

type CourseRepository struct{}

func (repo *CourseRepository) Save(ctx context.Context, course *courses.Course) error {
	return nil
}

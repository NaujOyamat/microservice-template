package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"Github.com/NaujOyamat/microservice-template/internal/domain/courses"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db}
}

func (repo *CourseRepository) Save(ctx context.Context, course *courses.Course) error {
	sqlCommand := "INSERT INTO courses (id, name, duration) VALUES ($1, $2, $3)"

	_, err := repo.db.ExecContext(ctx, sqlCommand, course.ID().String(), course.Name(), course.Duration())
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

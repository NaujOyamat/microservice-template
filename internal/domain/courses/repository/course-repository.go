package repository

import (
	"context"

	"Github.com/NaujOyamat/microservice-template/internal/domain/courses"
)

type ICourseRepository interface {
	Save(ctx context.Context, course *courses.Course) error
}

//go:generate mockery --name=ICourseRepository --outpkg=mocks --output=../../../platform/storage/mocks

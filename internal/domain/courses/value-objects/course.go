package valueobjects

import (
	"Github.com/NaujOyamat/microservice-template/internal/crosscutting"
	"Github.com/NaujOyamat/microservice-template/internal/domain/courses/errors"
)

type CourseID struct {
	value string
}

func NewCourseID(value string) (CourseID, error) {
	if !new(crosscutting.UUID).IsValid(value) {
		return CourseID{}, errors.ErrInvalidCourseID
	}
	return CourseID{value}, nil
}

func (c CourseID) String() string {
	return c.value
}

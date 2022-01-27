package courses

import valueobjects "Github.com/NaujOyamat/microservice-template/internal/domain/courses/value-objects"

type Course struct {
	id       valueobjects.CourseID
	name     string
	duration string
}

func NewCourse(id, name, duration string) (*Course, error) {
	courseId, err := valueobjects.NewCourseID(id)
	if err != nil {
		return nil, err
	}

	return &Course{
		id:       courseId,
		name:     name,
		duration: duration,
	}, nil
}

func (c *Course) ID() valueobjects.CourseID {
	return c.id
}

func (c *Course) Name() string {
	return c.name
}

func (c *Course) Duration() string {
	return c.duration
}

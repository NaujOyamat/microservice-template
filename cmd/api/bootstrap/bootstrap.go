package bootstrap

import (
	"Github.com/NaujOyamat/microservice-template/internal/platform/server"
	"Github.com/NaujOyamat/microservice-template/internal/platform/storage/postgres"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	dependencies := server.Dependencies{
		CourseRepo: &postgres.CourseRepository{},
	}

	srv := server.New(host, port, dependencies)
	return srv.Run()
}

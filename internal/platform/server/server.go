package server

import (
	"fmt"
	"log"

	courseRepo "Github.com/NaujOyamat/microservice-template/internal/domain/courses/repository"
	"Github.com/NaujOyamat/microservice-template/internal/platform/server/handlers/courses"
	"Github.com/NaujOyamat/microservice-template/internal/platform/server/handlers/health"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	httpAddr     string
	engine       *gin.Engine
	dependencies Dependencies
}

type Dependencies struct {
	CourseRepo courseRepo.ICourseRepository
}

func New(host string, port int, dependencies Dependencies) HttpServer {
	svr := HttpServer{
		httpAddr:     fmt.Sprintf("%s:%d", host, port),
		engine:       gin.New(),
		dependencies: dependencies,
	}

	svr.registerRoutes()
	return svr
}

func (s *HttpServer) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run()
}

func (s *HttpServer) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.dependencies.CourseRepo))
}

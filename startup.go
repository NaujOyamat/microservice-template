package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/golobby/config/v2"
	"github.com/golobby/container/pkg/container"
)

// Etructura encargada de la configuración
// del servicio web
type Startup struct{}

// Configura los servicios en el contenedor
// de dependencias del host
func (s *Startup) ConfigureServices(services *container.Container, configuration *config.Config) {}

// Configura las rutas del api
func (s *Startup) Configure(router *chi.Mux, configuration *config.Config) {}

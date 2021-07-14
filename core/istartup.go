package core

import (
	"github.com/go-chi/chi/v5"
	"github.com/golobby/config/v2"
	"github.com/golobby/container/pkg/container"
)

// Define las funciones que debe tener una
// estructura de inicio y configuración
type IStartup interface {
	// Configura los servicios en el contenedor
	// de dependencias del host
	ConfigureServices(*container.Container, *config.Config)
	// Configura las rutas del api
	Configure(*chi.Mux, *config.Config)
}

package core

import (
	"github.com/go-chi/chi/v5"
	"github.com/golobby/config/v2"
)

// Define las funciones que debe tener una
// estructura de inicio y configuración
type IStartup interface {
	// Configura las rutas y demás en el api
	ConfigureService(*chi.Mux, *config.Config)
}

package main

import (
	"Github.com/NaujOyamat/microservice-template/controllers"
	"Github.com/NaujOyamat/microservice-template/core"
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
func (s *Startup) Configure(router *chi.Mux, configuration *config.Config) {
	// Aplicamos los middleware por defecto
	core.SetDefaultMiddlewares(router)

	// Aqui se registran los controladores
	core.RegisterController("Greetting", func() interface{} {
		return new(controllers.GreettingController)
	})
}

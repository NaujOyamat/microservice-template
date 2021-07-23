package main

import (
	"Github.com/NaujOyamat/microservice-template/controllers"
	"Github.com/NaujOyamat/microservice-template/core"
	"github.com/go-chi/chi/v5"
	"github.com/golobby/config/v2"
)

// Etructura encargada de la configuración
// del servicio web
type Startup struct{}

// Configura las rutas y demás en el api
func (s *Startup) ConfigureService(router *chi.Mux, configuration *config.Config) {
	// Aplicamos los middleware por defecto
	core.SetDefaultMiddlewares(router)

	// Aqui se registran los controladores
	core.RegisterController("Greetting", func() interface{} {
		return new(controllers.GreettingController)
	})
}

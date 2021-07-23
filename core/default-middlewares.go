package core

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Asigna y configura los middleware por defecto
func SetDefaultMiddlewares(router *chi.Mux) {
	// Habilitamos cors
	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// Otros
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	router.Use(middleware.NoCache)
	router.Use(jsonContentType)
	// Timeout
	router.Use(middleware.Timeout(3600 * time.Second))

	// Agregamos el manejador de peticiones de controladores
	router.HandleFunc(fmt.Sprintf("/%s*", strings.ReplaceAll(ApiUrlPrefix, "\\", "")), RequestControllerHandler)
	// Agregamos manejador para archivos estáticos
	router.Handle("/*", http.FileServer(http.Dir("wwwroot")))
}

// Aplica el parametro contenType y lo asigna como application/json
func jsonContentType(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auxPrefix := fmt.Sprintf("/%s", strings.ReplaceAll(ApiUrlPrefix, "\\", ""))
		if len(r.URL.Path) >= 5 && strings.ToLower(string(r.URL.Path[:len(auxPrefix)])) == auxPrefix {
			w.Header().Add("Content-Type", "application/json")
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

package core

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golobby/config/v2"
	"github.com/golobby/config/v2/feeder"
)

// Instancia del host de la aplicación
var (
	Host *WebHost

	// Instancia del archivo de configuración
	Configuration *config.Config
)

// Encapsula la lógica necesaria para
// crear y levantar todo un entorno de servidor web
// proveiendo unas características básicas, como IoC
// lector del archivo appsettings.json y más...
type WebHost struct {
	port          string
	configuration *config.Config
	server        *chi.Mux
}

// Inicia la ejecución del host en el puerto
// especificado
func (h *WebHost) Run() {
	http.ListenAndServe(":"+h.port, h.server)
}

// Construye el host de la aplicación
func BuildWebHost(args []string, startupFunc StartupFunc) *WebHost {
	// Obtenemos el puerto de los argumentos entrantes si existe,
	// de lo contrario se usa el puerto por defecto 3001
	port, _ := GetPortArg(args, 3001)
	// Obtenemos el prefijo de los argumentos entrantes si existe,
	// de lo contrario se usa el api por defecto
	ApiUrlPrefix, _ = GetPrefixArg(args, "api")

	appsettings, err := config.New(feeder.Json{Path: "appsettings.json"})
	if err != nil {
		panic("No se ha podido cargar la configuración del archivo appsettings.json")
	}

	// Instanciamos el WebHost con el contenedor
	// de dependencias
	Host = &WebHost{
		port:          port,
		configuration: appsettings,
		server:        chi.NewRouter(),
	}

	Configuration = Host.configuration
	// Construimos y configuramos el startup
	startup := startupFunc()
	startup.ConfigureService(Host.server, Configuration)

	return Host
}

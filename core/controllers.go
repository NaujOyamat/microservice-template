package core

import (
	"strings"
)

// Prefijo usado en la url que apunta a los controladores
var (
	ApiUrlPrefix string = `api\/`

	// Contenedor de controladores registrados
	ctrlContainer map[string]resolverFunc
)

// Registra un controlador con una clave
func RegisterController(key string, resolver resolverFunc) {
	if ctrlContainer == nil {
		ctrlContainer = make(map[string]resolverFunc)
	}

	key = strings.ToLower(strings.Trim(key, ""))
	if len(key) == 0 {
		panic("RegisterController: La clave de registro de un controlador no puede ser vacía")
	}

	if resolver == nil {
		panic("RegisterController: La función resolvedora de un controlador no puede ser nula")
	}

	ctrlContainer[key] = resolver
}

// Obtiene una instancia del controlador registrado con clave key
func GetControllerInstance(key string) (interface{}, bool) {
	resolver, ok := ctrlContainer[key]
	if !ok {
		return nil, false
	}
	return resolver(), true
}

package core

// Función Startup
type StartupFunc func() IStartup

// Función resolvedora de controladores
type resolverFunc func() interface{}

// Encapsula los datos de una respuesta
// a una solicitud de acción en un controlador
type RequestResult struct {
	IsSuccessful bool        `json:"isSuccessful"`
	IsError      bool        `json:"isError"`
	ErrorMessage string      `json:"errorMessage"`
	Messages     []string    `json:"messages"`
	Result       interface{} `json:"result"`
}

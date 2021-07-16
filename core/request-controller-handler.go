package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Manejador de peticiones al api. Realiza la lógica necesaria
// para invocar el controladore correspondiente según la petición
func RequestControllerHandler(w http.ResponseWriter, r *http.Request) {
	//Obtenemos el nombre del controlador y la acción
	re := regexp.MustCompile(fmt.Sprintf(`^\/%s([a-zA-Z_][a-zA-Z0-9_]+)\/([a-zA-Z_][a-zA-Z0-9_]+)`, ApiUrlPrefix))
	matches := re.FindAllStringSubmatch(r.URL.Path, 1)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		controllerName := matches[0][1]
		// Buscamos el controlador en el contenedor
		controllerResolver, ok := ctrlContainer[strings.ToLower(controllerName)]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			// Buscamos la acción entre todos los metodos del controlador
			actionName := matches[0][2]
			// Obtenemos la instancia del controlador
			controllerInstance := controllerResolver()
			// Obtenemos su tipo por reflection
			controllerType := reflect.TypeOf(controllerInstance)
			// Iteramos cada uno de sus metodos en busca de uno que coincida con la acción y el verbo http
			for i := 0; i < controllerType.NumMethod(); i++ {
				// Obtenemos el tipo del metodo
				method := controllerType.Method(i)
				// Validamos por el nombre y el verbo http si corresponde a la acción
				if strings.EqualFold(method.Name, fmt.Sprintf("%s%s", r.Method, actionName)) {
					// Validamos que el método retorne solo un parámetro y que reciba solo un parámetro,
					// además de que ese parámetro sea de tipo Struct
					if method.Type.NumOut() != 1 {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("ERROR: El método no puede retornar más de un parámetro"))
						return
					}
					if method.Type.NumIn() != 2 || (method.Type.In(1).Kind() != reflect.Struct && method.Type.In(1).Kind() != reflect.Interface) {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("ERROR: El método solo puede tener un parámetro de tipo Struct"))
						return
					}
					// Obtenemos el valor del controlador por reflection
					controllerValue := reflect.ValueOf(controllerInstance)
					// Ejecutamos el método del controlador pasando los parámetros a partir de los parámetros de la petición
					methodResult, err := executeControllerAction(&controllerValue, &method, r)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("ERROR: Serializando la respuesta a formato JSON"))
					} else {
						// Serializamos el resultado a json
						jsonResult, err := json.Marshal(methodResult[0].Interface())
						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
							w.Write([]byte("ERROR: Serializando la respuesta a formato JSON"))
						} else {
							// Retornamos la respuesta JSON
							w.WriteHeader(http.StatusOK)
							w.Write(jsonResult)
						}
					}
				}
			}
			// No se encontro ningún metodo que haga match con la acción
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// Ejecuta la acción en el controlador según el verbo http
// leyendo los parámetros del queryString o el cuerpo de la petición
func executeControllerAction(controllerValue *reflect.Value, method *reflect.Method, request *http.Request) ([]reflect.Value, error) {
	var result []reflect.Value = nil
	err := fmt.Errorf("executeControllerAction: Verbo http no soportado (%v)", request.Method)
	// Si el verbo http es Get o Delete, los parametros se toman del queryString
	if request.Method == http.MethodGet || request.Method == http.MethodDelete {
		result = make([]reflect.Value, 0)
		err = nil
		// Creamos la instancia de la estructura
		pIn := reflect.New(method.Type.In(1))
		// Recorremos los parametros del queryString y los mapeamos a la estructura de entrada
		for k, v := range request.URL.Query() {
			// Validamos si el parametro existe en la estructura
			if pIn.Elem().FieldByName(k).IsValid() && len(v) > 0 {
				// Mapeamos el parámetro según el tipo de dato del campo en la estructura
				mapParamToField(pIn.Elem().FieldByName(k), v)
			} else {
				result = nil
				err = fmt.Errorf("executeControllerAction: Parametro (%v) inválido", k)
				break
			}
		}
		// Si no hay error, la estructura se mapeo con exito
		if err == nil {
			result = controllerValue.MethodByName(method.Name).Call([]reflect.Value{pIn.Elem()})
		}
	} else if request.Method == http.MethodPost || request.Method == http.MethodPut {
		defer request.Body.Close()
		// Si el verbo http es Post o Put, los parametros se toman del cuerpo de la petición
		err = nil

		// Creamos la instancia de la estructura
		target := reflect.New(method.Type.In(1)).Interface()
		// Deserializamos el cuerpo
		err := json.NewDecoder(request.Body).Decode(&target)
		if err == nil {
			result = controllerValue.MethodByName(method.Name).Call([]reflect.Value{reflect.ValueOf(target).Elem()})
		}
	}

	return result, err
}

// Mapea el campo con el parámetro según su tipo de dato
func mapParamToField(field reflect.Value, param []string) {
	fieldValue := field.Interface()
	switch fieldValue.(type) {
	case string:
		field.SetString(param[0])
	case time.Time:
		t, err := time.Parse(time.RFC3339, param[0])
		if err == nil {
			field.Set(reflect.ValueOf(t))
		}
	case int8:
		vInt, err := strconv.ParseInt(param[0], 10, 8)
		if err == nil {
			field.SetInt(vInt)
		}
	case int16:
		vInt, err := strconv.ParseInt(param[0], 10, 16)
		if err == nil {
			field.SetInt(vInt)
		}
	case int32:
		vInt, err := strconv.ParseInt(param[0], 10, 32)
		if err == nil {
			field.SetInt(vInt)
		}
	case int64:
	case int:
		vInt, err := strconv.ParseInt(param[0], 10, 64)
		if err == nil {
			field.SetInt(vInt)
		}
	case bool:
		vBool, err := strconv.ParseBool(param[0])
		if err == nil {
			field.SetBool(vBool)
		}
	case float32:
		vFloat, err := strconv.ParseFloat(param[0], 32)
		if err == nil {
			field.SetFloat(vFloat)
		}
	case float64:
		vFloat, err := strconv.ParseFloat(param[0], 64)
		if err == nil {
			field.SetFloat(vFloat)
		}
	}
}

package controllers

import (
	"Github.com/NaujOyamat/microservice-template/models"
)

// Controlador de saludos
type GreettingController struct{}

// Obtiene un saludo al mundo
func (*GreettingController) GetHelloWorld(p models.Person) models.Person {
	return p //fmt.Sprintf("Hola Mundo, %d - %s!", p.Id, p.Name)
}

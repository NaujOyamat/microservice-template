package controllers

import (
	"Github.com/NaujOyamat/microservice-template/models"
)

// Controlador de saludos
type GreettingController struct{}

// Obtiene una persona por la url y la retorna
func (*GreettingController) GetPersonWithGet(p models.Person) models.Person {
	return p
}

// Obtiene una persona por el cuerpo de la petición y la retorna
func (*GreettingController) PostPersonWithPost(p models.Person) models.Person {
	return p
}

package main

import (
	"api-rest-go/models"
	"api-rest-go/routes"
	"fmt"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: int(1), Name: string("Tronxo"), Biography: string("Tronxo demais")},
	}
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}

package main

import (
	"fmt"

	"github.com/jeantarlles/go-rest-api/database"
	"github.com/jeantarlles/go-rest-api/models"
	"github.com/jeantarlles/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor rest com Go")
	routes.HandleRequest()
}

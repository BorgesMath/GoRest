package main

import (
	"APIGoRest/database"
	"APIGoRest/models"
	"APIGoRest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectarComBanco()
	fmt.Println("Iniciando o servidor Rest go")
	routes.HandleRequest()
}

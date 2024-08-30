package main

import (
	"APIGoRest/database"
	"APIGoRest/routes"
	"fmt"
)

func main() {
	database.ConectarComBanco()
	fmt.Println("Iniciando o servidor Rest go")

	routes.HandleRequest()
}

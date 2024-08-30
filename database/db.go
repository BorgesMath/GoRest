package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declara a variável global DB e err
var (
	DB  *gorm.DB
	err error
)

func ConectarComBanco() {
	// Define a string de conexão
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	// Abre a conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})

	if err != nil {
		// Em caso de erro, registra o erro e encerra a aplicação
		log.Panic("Erro ao conectar com o DB: ", err)
	}

	// Você pode adicionar aqui outras configurações ou inicializações necessárias
}

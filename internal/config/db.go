package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("DB_SSLMODE")

	//DSN = Data Source Name → string de conexão no formato esperado pelo driver do PostgreSQL + GORM
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, sslMode)

	//Abre a conexão com GORM usando o driver do PostgreSQL.
	database, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Erro ao conectar com o banco de dados:", err)
	}

	DB = database
	fmt.Println("✅ Banco de dados conectado com sucesso")
}

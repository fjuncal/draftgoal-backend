package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbUser := GetEnv("POSTGRES_USER", "postgres")
	dbPassword := GetEnv("POSTGRES_PASSWORD", "postgres")
	dbName := GetEnv("POSTGRES_DB", "draftgoal")
	sslMode := GetEnv("DB_SSLMODE", "disable")

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

	return DB
}

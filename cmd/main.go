package main

import (
	"draftgoal-backend/internal/config"
	"draftgoal-backend/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//caregando o arquivo .env
	config.LoadEnv()

	//Conecta com o banco de dados
	db := config.ConnectDatabase()

	// Configura as rotas e injeções
	router.Initialize(db)

	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

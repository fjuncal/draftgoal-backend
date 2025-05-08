package main

import (
	"draftgoal-backend/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//caregando o arquivo .env
	config.LoadEnv()

	//Conecta com o banco de dados
	config.ConnectDatabase()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "DraftGoal is running!")
	})
	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

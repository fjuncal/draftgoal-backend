package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnv Funcao que carrega o arquivo .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Erro ao carregar .env")
	}
}

// GetEnv Funcao que retorna o valor de uma variavel de ambiente ou um fallback
func GetEnv(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

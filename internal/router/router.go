package router

import (
	"draftgoal-backend/internal/handler"
	"draftgoal-backend/internal/repository"
	"draftgoal-backend/internal/service"
	"gorm.io/gorm"
	"net/http"
)

// Initialize define as rotas e injeções manuais
func Initialize(db *gorm.DB) {
	// Injeta as dependências
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Rotas
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("DraftGoal is running!"))
		if err != nil {
			http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userHandler.CreateUserHandler(w, r)
			return
		}
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	})
}

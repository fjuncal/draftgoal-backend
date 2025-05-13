package handler

import (
	"draftgoal-backend/internal/model"
	"draftgoal-backend/internal/service"
	"encoding/json"
	"net/http"
)

// UserHandler representa o controlador HTTP para usuários
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler cria uma nova instância de UserHandler
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUserHandler lida com o endpoint POST /users
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decodifica o body JSON para struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Chama a regra de negócio
	if err := h.userService.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Remove a senha antes de retornar
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

package service

import (
	"draftgoal-backend/internal/model"
	"draftgoal-backend/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// UserService define os métodos relacionados à lógica de negócio do usuário
type UserService interface {
	// CreateUser cria um novo usuário após validação e criptografia da senha
	CreateUser(user *model.User) error
}

// userService é a implementação de UserService que usa o repositório de usuários
type userService struct {
	repo repository.UserRepository
}

// NewUserService cria uma nova instância de userService com o repositório injetado
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser valida, criptografa e persiste um novo usuário
func (s *userService) CreateUser(user *model.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("e-mail já cadastrado")
	}

	hashed, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	return s.repo.Create(user)
}

// hashPassword gera o hash seguro da senha do usuário usando bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

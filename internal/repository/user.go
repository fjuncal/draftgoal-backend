package repository

import (
	"draftgoal-backend/internal/model"
	"gorm.io/gorm"
)

// UserRepository define a interface para a persistencia de usuarios
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

// userRepository é a implementação concreta
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository injeta o *gorm.DB e retorna a interface
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create insere um novo usuário no banco
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// FindByEmail busca um usuário pelo email
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

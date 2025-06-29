package services

import (
	"errors"

	"golang/internal/models"

	"gorm.io/gorm"
)

// UserService gerencia operações relacionadas a usuários.
type UserService struct {
	db *gorm.DB
}

// NewUserService cria uma nova instância do UserService.
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser cria um novo usuário.
func (s *UserService) CreateUser(user *models.User) error {
	// Verificar se o email já existe
	var existingUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already exists")
	}

	// TODO: Hash da senha antes de salvar
	// user.Password = hashPassword(user.Password)

	return s.db.Create(user).Error
}

// GetUserByID busca um usuário pelo ID.
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail busca um usuário pelo email.
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser atualiza um usuário.
func (s *UserService) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

// DeleteUser remove um usuário (soft delete).
func (s *UserService) DeleteUser(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

// ListUsers lista todos os usuários com paginação.
func (s *UserService) ListUsers(offset, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Contar total de registros
	if err := s.db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Buscar usuários com paginação
	if err := s.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

package models

import (
	"time"

	"gorm.io/gorm"
)

// User representa um usuário no sistema
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Name      string         `json:"name" gorm:"not null"`
	Password  string         `json:"-" gorm:"not null"` // "-" oculta o campo no JSON
	Active    bool           `json:"active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName especifica o nome da tabela
func (User) TableName() string {
	return "users"
}

// BeforeCreate hook executado antes de criar um usuário
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook executado antes de atualizar um usuário
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

package userService

import (
	"gorm.io/gorm"
	"pet_project/internal/taskService"
)

type Users struct {
	gorm.Model
	Email    string                `json:"email"`
	Password string                `json:"password"`
	Tasks    []taskService.Message `json:"tasks" gorm:"foreignKey:UserID"`
}

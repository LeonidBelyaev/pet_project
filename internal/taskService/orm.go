package taskService

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task   string `json:"task"`    // Наш сервер будет ожидать json c полем text
	IsDone bool   `json:"is_done"` // В GO используем CamelCase, в Json - snake
	UserID uint   `json:"user_id"`
}

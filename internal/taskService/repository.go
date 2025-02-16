package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Message) (Message, error)
	GetAllTasks() ([]Message, error)
	UpdateTaskById(id uint, task Message) (Message, error)
	DeleteTaskById(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Message) (Message, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Message, error) {
	var tasks []Message
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskById(id uint, task Message) (Message, error) {
	var updateTask Message
	if err := r.db.First(&updateTask, id).Error; err != nil {
		return updateTask, err
	}
	if task.Task != "" {
		updateTask.Task = task.Task
	}
	if task.IsDone {
		updateTask.IsDone = task.IsDone
	}
	if err := r.db.Save(&updateTask).Error; err != nil {
		return updateTask, err
	}
	return updateTask, nil
}

func (r *taskRepository) DeleteTaskById(id uint) error {
	if err := r.db.Delete(&Message{}, id).Error; err != nil {
		return err
	}
	return nil
}

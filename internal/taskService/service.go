package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Message) (Message, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTask() ([]Message, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskById(id uint, task Message) (Message, error) {
	return s.repo.UpdateTaskById(id, task)
}

func (s *TaskService) DeleteTaskById(id uint) error {
	return s.repo.DeleteTaskById(id)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]Message, error) {
	return s.repo.GetTasksByUserID(userID)
}

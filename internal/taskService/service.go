package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(userID uint, task Message) (Message, error) {
	task.UserID = userID
	return s.repo.CreateTask(userID, task)
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

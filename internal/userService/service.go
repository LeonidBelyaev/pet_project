package userService

import "pet_project/internal/taskService"

type UsersService struct {
	repo UsersRepository
}

func NewUsersService(repo UsersRepository) *UsersService {
	return &UsersService{repo: repo}
}

func (u *UsersService) CreateUser(user Users) (Users, error) {
	return u.repo.CreateUser(user)
}

func (u *UsersService) GetAllUsers() ([]Users, error) {
	return u.repo.GetAllUsers()
}

func (u *UsersService) UpdateUserById(id uint, user Users) (Users, error) {
	return u.repo.UpdateUserById(id, user)
}

func (u *UsersService) DeleteUserById(id uint) error {
	return u.repo.DeleteUserById(id)
}

func (u *UsersService) GetTasksForUser(userID uint) ([]taskService.Message, error) {
	return u.repo.GetTasksForUser(userID)
}

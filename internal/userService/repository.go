package userService

import "gorm.io/gorm"

type UsersRepository interface {
	CreateUser(user Users) (Users, error)
	GetAllUsers() ([]Users, error)
	UpdateUserById(id uint, user Users) (Users, error)
	DeleteUserById(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user Users) (Users, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return Users{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserById(id uint, user Users) (Users, error) {
	var updateUser Users
	if err := r.db.First(&updateUser, id).Error; err != nil {
		return updateUser, err
	}
	if user.Email != "" {
		updateUser.Email = user.Email
	}
	if user.Password != "" {
		updateUser.Password = user.Password
	}
	if err := r.db.Save(&updateUser).Error; err != nil {
		return updateUser, err
	}
	return updateUser, nil
}

func (r *userRepository) DeleteUserById(id uint) error {
	if err := r.db.Delete(&Users{}, id).Error; err != nil {
		return err
	}
	return nil
}

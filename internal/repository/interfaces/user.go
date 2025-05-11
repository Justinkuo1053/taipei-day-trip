package interfaces

import "taipei-day-trip/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type UserService interface {
	GetUserByID(id int) (*models.User, error)
}

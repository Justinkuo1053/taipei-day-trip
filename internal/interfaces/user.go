package interfaces

import "taipei-day-trip-go-go/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
}

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (string, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

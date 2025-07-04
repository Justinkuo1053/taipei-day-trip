package services

import (
	"errors"
	"fmt"
	"time"

	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo      interfaces.UserRepository
	JwtSecret string
}

func NewUserService(repo interfaces.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		Repo:      repo,
		JwtSecret: jwtSecret,
	}
}

func (s *UserService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密碼加密失敗: %w", err)
	}
	user.Password = string(hashedPassword)
	if err := s.Repo.Create(user); err != nil {
		return fmt.Errorf("註冊失敗: %w", err)
	}
	return nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("帳號或密碼錯誤")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("帳號或密碼錯誤")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("生成 token 失敗: %w", err)
	}
	return tokenString, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.Repo.GetByEmail(email)
}

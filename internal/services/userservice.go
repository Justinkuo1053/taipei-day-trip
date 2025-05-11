package services

import (
	"errors"
	"fmt"
	"time"

	"taipei-day-trip/internal/models"     // 修正 import 路徑
	"taipei-day-trip/internal/repository" // 修正 import 路徑
	"taipei-day-trip/internal/repository/interfaces"

	"github.com/golang-jwt/jwt/v4" // 添加 JWT 套件
	"golang.org/x/crypto/bcrypt"   // 添加 bcrypt 套件
)

type userService struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewUserService(repo repository.UserRepository, jwtSecret string) interfaces.UserService {
	return &userService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *userService) Register(user *models.User) error {
	// 加密密碼
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密碼加密失敗: %w", err)
	}
	user.Password = string(hashedPassword)

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("註冊失敗: %w", err)
	}
	return nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", fmt.Errorf("登入失敗: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("密碼錯誤")
	}

	// 生成 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("生成 token 失敗: %w", err)
	}

	return tokenString, nil
}

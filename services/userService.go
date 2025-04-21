package services

import (
	"log"

	"github.com/doneowth/interfaces"
	"github.com/doneowth/models"
	"gorm.io/gorm"
)

type userService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) interfaces.UserService {
	return &userService{DB: db}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	if err := s.DB.Create(user).Error; err != nil {
		log.Fatalln("failed to create user")
		return nil, err
	}
	return user, nil
}

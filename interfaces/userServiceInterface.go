package interfaces

import "github.com/doneowth/models"

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
}

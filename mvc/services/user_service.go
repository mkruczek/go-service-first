package services

import (
	"github.com/mkruczek/go-service-first/mvc/domain"
	"github.com/mkruczek/go-service-first/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (us *userService) GetUsers() []*domain.User {
	return domain.UserDao.GetUsers()
}

func (us *userService) GetUser(id uint64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(id)
}

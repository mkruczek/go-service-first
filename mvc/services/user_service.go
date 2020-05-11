package services

import (
	"github.com/mkruczek/go-service-first/mvc/domain"
	"github.com/mkruczek/go-service-first/mvc/utils"
)

func GetUsers() []*domain.User {
	return domain.GetUsers()
}

func GetUser(id uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(id)
}

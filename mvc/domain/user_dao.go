package domain

import (
	"fmt"
	"net/http"

	"github.com/mkruczek/go-service-first/mvc/utils"
)

var (
	UserDao userDaoInterface
	//mock db
	users = map[uint64]*User{
		1: {ID: 1, Firstname: "Foo", Lastname: "Bar", Email: "foo@bar.io"},
		2: {ID: 2, Firstname: "Fabio", Lastname: "Ferrari", Email: "Fabio@Ferrari.io"},
	}
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(uint64) (*User, *utils.ApplicationError)
	GetUsers() []*User
}

type userDao struct{}

func (ud *userDao) GetUsers() []*User {
	result := []*User{}

	for _, v := range users {
		result = append(result, v)
	}

	return result
}

func (ud *userDao) GetUser(id uint64) (*User, *utils.ApplicationError) {
	if user := users[id]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Msg:    fmt.Sprintf("User not faund: %d", id),
		Status: http.StatusNotFound,
		Code:   "not_faund"}
}

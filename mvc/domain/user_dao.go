package domain

import (
	"fmt"
	"net/http"

	"github.com/mkruczek/go-service-first/mvc/utils"
)

var (
	//mock db
	users = map[uint64]*User{
		1: {ID: 1, Firstname: "Foo", Lastname: "Bar", Email: "foo@bar.io"},
		2: {ID: 2, Firstname: "Fabio", Lastname: "Ferrari", Email: "Fabio@Ferrari.io"},
	}
)

func GetUsers() []*User {
	result := make([]*User, len(users))

	for _, v := range users {
		result = append(result, v)
	}

	return result
}

func GetUser(id uint64) (*User, *utils.ApplicationError) {
	if user := users[id]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Msg:    fmt.Sprintf("User not faund: %d", id),
		Status: http.StatusNotFound,
		Code:   "not_faund"}
}

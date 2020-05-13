package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mkruczek/go-service-first/mvc/domain"
	"github.com/mkruczek/go-service-first/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	udm          userDaoMock
	getUserFunc  func(id uint64) (*domain.User, *utils.ApplicationError)
	getUsersFunc func() []*domain.User
)

func init() {
	domain.UserDao = &userDaoMock{}
}

type userDaoMock struct {
}

func (udm *userDaoMock) GetUser(id uint64) (*domain.User, *utils.ApplicationError) {
	return getUserFunc(id)
}

func (udm *userDaoMock) GetUsers() []*domain.User {
	return getUsersFunc()
}

func TestGetUserNotFound(t *testing.T) {

	getUserFunc = func(id uint64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Status: http.StatusNotFound,
			Msg:    fmt.Sprintf("User not faund: %d", id),
			Code:   "not_faund"}
	}

	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_faund", err.Code)
	assert.EqualValues(t, "User not faund: 0", err.Msg)
}

func TestGetUserFound(t *testing.T) {

	getUserFunc = func(id uint64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{ID: 418, Firstname: "Mock1", Lastname: "Mock2", Email: "mock@mock.com"}, nil
	}

	user, err := UserService.GetUser(0)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 418, user.ID)
	assert.EqualValues(t, "Mock1", user.Firstname)
	assert.EqualValues(t, "Mock2", user.Lastname)
	assert.EqualValues(t, "mock@mock.com", user.Email)
}

func TestGetUsers(t *testing.T) {
	getUsersFunc = func() []*domain.User {
		return []*domain.User{
			{ID: 1, Firstname: "first", Email: "first@first.com"},
			{ID: 2, Firstname: "second", Email: "second@second.com"},
		}
	}

	result := UserService.GetUsers()

	assert.EqualValues(t, 2, len(result))
	assert.EqualValues(t, 2, cap(result))

	assert.EqualValues(t, 1, result[0].ID)
	assert.EqualValues(t, "first", result[0].Firstname)
	assert.EqualValues(t, "first@first.com", result[0].Email)
	assert.EqualValues(t, 2, result[1].ID)
	assert.EqualValues(t, "second", result[1].Firstname)
	assert.EqualValues(t, "second@second.com", result[1].Email)
}

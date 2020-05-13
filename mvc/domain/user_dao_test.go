package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	users := UserDao.GetUsers()

	assert.EqualValues(t, 2, len(users), "expect list with two users")

	// if len(users) != 2 {
	// 	t.Error("expect list with two users")
	// }
}

func TestGetUserNonExist(t *testing.T) {
	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "we don't expect user witch id 0")
	assert.NotNil(t, err, "we don't expect user witch id 0")
	assert.EqualValues(t, http.StatusNotFound, err.Status)

	// if user != nil {
	// 	t.Error("we don't expect user witch id 0")
	// }

	// if err == nil {
	// 	t.Error("we don't expect user witch id 0")
	// }

	// if err.Status != http.StatusNotFound {
	// 	t.Error("expect 404 status")
	// }
}

func TestGetUserExist(t *testing.T) {

	user, err := UserDao.GetUser(1)

	assert.Nil(t, err, "expect nil err")
	assert.EqualValues(t, 1, user.ID, "expect ID 1")
	assert.EqualValues(t, "Foo", user.Firstname, "expect  Firstname foo")
	assert.EqualValues(t, "Bar", user.Lastname, "expect Lastname bar")
	assert.EqualValues(t, "foo@bar.io", user.Email, "expect email foo@bar.io")
}

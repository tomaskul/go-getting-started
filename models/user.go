package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type UserPersistence struct {
}

var (
	users  []*User
	nextID = 1
)

func (up UserPersistence) GetUsers() []*User {
	return users
}

func (up UserPersistence) AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new User must not include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func (up UserPersistence) GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with ID: '%v' not found", id)
}

func (up UserPersistence) UpdateUser(u User) (User, error) {
	if u.ID <= 0 {
		return User{}, fmt.Errorf("supplied User with invalid ID: '%v'", u.ID)
	}

	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with ID: '%v' not found", u.ID)
}

func (up UserPersistence) RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID: '%v' not found", id)
}

func NewUserPersistence() *UserPersistence {
	return &UserPersistence{}
}

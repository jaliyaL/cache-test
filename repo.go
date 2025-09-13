package main

import (
	"errors"
	"fmt"
)

// Sentinel error
var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID   int
	Name string
}

type UserRepo struct {
	data map[int]User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		data: map[int]User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}
}

func (r *UserRepo) GetUser(id int) (User, error) {
	user, ok := r.data[id]
	if !ok {
		return User{}, fmt.Errorf("repo: id=%d: %w", id, ErrUserNotFound)
	}
	return user, nil
}

package main

import (
	"errors"
	"fmt"
)

type UserService struct {
	repo *UserRepo
}

func NewUserService(repo *UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(id int) (User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		// Add context and re-wrap
		if errors.Is(err, ErrUserNotFound) {
			return User{}, fmt.Errorf("service: user %d does not exist: %w", id, err)
		}
		return User{}, fmt.Errorf("service: failed fetching user %d: %w", id, err)
	}
	return user, nil
}

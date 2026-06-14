package service

import (
	"github.com/khazeez/user-service/internal/domain"
	"github.com/khazeez/user-service/internal/repository"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(
    repo repository.UserRepository,
) *UserService {

    return &UserService{
        repo: repo,
    }
}

func (s *UserService) CreateUser(
    name string,
    email string,
) (*domain.User, error) {

    user := &domain.User{
        Name: name,
        Email: email,
    }

    err := s.repo.Create(user)

    if err != nil {
        return nil, err
    }

    return user, nil
}

func (s *UserService) GetUser(
    id int64,
) (*domain.User, error) {

    return s.repo.GetByID(id)
}

func (s *UserService) ListUsers() (
    []domain.User,
    error,
) {
    return s.repo.List()
}

func (s *UserService) DeleteUser(
    id int64,
) error {

    return s.repo.Delete(id)
}
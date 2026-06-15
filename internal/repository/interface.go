package repository

import "github.com/khazeez/user-service/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
	List() ([]domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
}

package repository

import (
	"errors"
	"github.com/khazeez/user-service/internal/domain"
)

type MemoryRepo struct {
	users  map[int64]domain.User
	nextID int64
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		users:  make(map[int64]domain.User),
		nextID: 1,
	}
}

func (r *MemoryRepo) Create(
	user *domain.User,
) error {

	user.ID = r.nextID

	r.users[user.ID] = *user

	r.nextID++

	return nil
}

func (r *MemoryRepo) GetByID(
	id int64,
) (*domain.User, error) {

	user, ok := r.users[id]

	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *MemoryRepo) List() (
	[]domain.User,
	error,
) {

	result := []domain.User{}

	for _, u := range r.users {
		result = append(result, u)
	}

	return result, nil
}

func (r *MemoryRepo) Update(
	user *domain.User,
) error {

	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}

	r.users[user.ID] = *user

	return nil
}

func (r *MemoryRepo) Delete(
	id int64,
) error {

	delete(r.users, id)

	return nil
}

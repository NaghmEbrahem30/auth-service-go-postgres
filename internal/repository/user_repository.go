package repository

import (
	"errors"
	"sync"

	"auth-service-go-postgres/internal/domain"
)

var ErrUserExists = errors.New("user already exists")

type UserRepository interface {
	Create(user domain.User) error
	FindByEmail(email string) (domain.User, bool)
}

type InMemoryUserRepository struct {
	mu      sync.RWMutex
	byEmail map[string]domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{byEmail: make(map[string]domain.User)}
}

func (r *InMemoryUserRepository) Create(user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.byEmail[user.Email]; ok {
		return ErrUserExists
	}
	r.byEmail[user.Email] = user
	return nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (domain.User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.byEmail[email]
	return u, ok
}

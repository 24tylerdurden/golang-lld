package repository

import (
	"rate-limiter/internal/models"
	"sync"
)

type InMemoryRepository struct {
	data  map[string]*models.UserData
	mutex sync.Mutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: make(map[string]*models.UserData),
	}
}

func (i *InMemoryRepository) Get(userId string) (*models.UserData, bool) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	user, exists := i.data[userId]
	return user, exists
}

func (i *InMemoryRepository) Set(userId string, data *models.UserData) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.data[userId] = data
}

package interfaces

import (
	"rate-limiter/internal/models"
	"rate-limiter/internal/repository"
	"sync"
	"time"
)

type FixedWindow struct {
	repo       repository.UserRepository
	limit      int
	mutex      sync.Mutex
	windowSize time.Duration
}

func NewFixedWindow(repo repository.UserRepository) *FixedWindow {
	return &FixedWindow{
		repo:       repo,
		limit:      5,
		windowSize: time.Minute,
	}
}

func (f *FixedWindow) Allow(userId string) bool {
	f.mutex.Lock()

	defer f.mutex.Unlock()

	// check whether we can allow or not

	now := time.Now()

	user, exists := f.repo.Get(userId)

	if !exists {
		f.repo.Set(userId, models.NewUserData())
		return true
	}

	if now.Sub(user.WindowStart) < f.windowSize {
		if user.Count < f.limit {
			user.Count++
			return true
		}
	}

	// User limit expired
	user.Count = 1
	user.WindowStart = time.Now().Add(time.Minute)
	return false
}

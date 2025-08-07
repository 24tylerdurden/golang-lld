package factory

import (
	"rate-limiter/internal/interfaces"
	"rate-limiter/internal/repository"
)

type LimiterType string

const (
	FixedWindow   LimiterType = "fixed_window"
	SlidingWindow LimiterType = "slidin_window"
)

func NewRateLimiterFactory(repo repository.UserRepository, limiterType LimiterType) interfaces.RateLimiter {
	switch limiterType {
	case FixedWindow:
		return interfaces.NewFixedWindow(repo)
	default:
		return interfaces.NewFixedWindow(repo)
	}
}

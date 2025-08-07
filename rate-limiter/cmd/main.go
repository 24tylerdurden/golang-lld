package main

import (
	"fmt"
	"rate-limiter/internal/factory"
	"rate-limiter/internal/repository"
	"time"
)

func main() {
	repo := repository.NewInMemoryRepository()
	rateLimiter := factory.NewRateLimiterFactory(repo, "fixed_window")

	userId := "user1"

	for i := 1; i <= 7; i++ {
		allowed := rateLimiter.Allow(userId)

		fmt.Printf("Request %d allowed: %v\n", i, allowed)
		time.Sleep(2 * time.Second)
	}
}

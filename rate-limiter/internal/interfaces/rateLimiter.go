package interfaces

type RateLimiter interface {
	Allow(userid string) bool
}

package repository

import (
	"rate-limiter/internal/models"
)

type UserRepository interface {
	Get(userId string) (*models.UserData, bool)
	Set(userId string, data *models.UserData)
}

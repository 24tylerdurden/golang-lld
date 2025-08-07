package models

import "time"

type UserData struct {
	Count       int
	WindowStart time.Time
}

func NewUserData() *UserData {
	now := time.Now().Add(time.Minute)
	return &UserData{
		Count:       1,
		WindowStart: now,
	}
}

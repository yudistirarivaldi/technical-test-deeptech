package model

import (
	"time"
)

type Users struct {
	ID          int64
	FirstName   string
	LastName    string
	Email       string
	Password    string
	DateOfBirth time.Time
	Gender      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

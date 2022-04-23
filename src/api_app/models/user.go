package models

import "time"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsActive  bool
	CreatedAt time.Time
}

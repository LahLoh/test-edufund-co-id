package model

import "time"

type User struct {
	ID        uint64
	Fullname  string
	Username  string
	Password  string
	CreatedAt time.Time
}

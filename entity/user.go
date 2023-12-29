package entity

import "time"

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Age      string    `json:"age"`
	Email    string    `json:"email"`
	Password string    `json:"pasword"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

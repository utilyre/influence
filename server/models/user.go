package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `json:"id" validate:"isdefault"`
	CreatedAt time.Time `json:"created_at" validate:"isdefault"`
	UpdatedAt time.Time `json:"updated_at" validate:"isdefault"`

	Name     string `json:"name" validate:"required,min=3" gorm:"not null"`
	Email    string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=256" gorm:"not null"`
}

func (u *User) Validate(exceptions ...string) error {
	validate := validator.New()
	return validate.StructExcept(u, exceptions...)
}

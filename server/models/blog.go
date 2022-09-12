package models

import (
	"errors"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Blog struct {
	ID        uint      `json:"id" validate:"isdefault"`
	CreatedAt time.Time `json:"created_at" validate:"isdefault"`
	UpdatedAt time.Time `json:"updated_at" validate:"isdefault"`

	AuthorID uint `json:"-" validate:"isdefault"`
	Author   User `json:"author" validate:"-"`

	Title   string `json:"title" validate:"required,min=8,max=64" gorm:"not null"`
	Content string `json:"content" validate:"required,min=16,max=1024" gorm:"not null"`
}

func (b *Blog) Validate(exceptions ...string) error {
	validate := validator.New()
	err := validate.StructExcept(b, exceptions...)
	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)
	translator := ut.New(en.New())
	translations := errs.Translate(translator.GetFallback())

	text := ""
	for _, tag := range translations {
		text += tag + "\n"
	}

	return errors.New(text)
}

func (b *Blog) Merge(blog Blog) {
	b.Title = blog.Title
	b.Content = blog.Content
}

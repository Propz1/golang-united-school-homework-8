package services

import (
	"io"
)

type UserService interface {
	Add(item string, w io.Writer) error
	GetUserById(id string, w io.Writer) error
	GetList(w io.Writer) error
	Remove(id string, w io.Writer) error
}

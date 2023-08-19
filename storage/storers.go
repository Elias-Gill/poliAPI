package storage

import "github.com/elias-gill/poliapi/types"

type UserStorer interface {
	GetById(string) (*types.User, error)
	UpdateSection(int) (*types.User, error)
	Delete(string) error
	Insert(types.User) error
}

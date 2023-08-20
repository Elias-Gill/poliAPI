package storage

import "github.com/elias-gill/poliapi/types"

type UserStorer interface {
	GetById(string) (*types.User, error)
	Update(string, string, any) error
	Delete(string) error
	Insert(types.User) error
}

package storage

import "github.com/elias-gill/poliapi/types"

type UserStorer interface {
	GetById(id string) (*types.User, error)
	Update(id string, newData *types.User) error
	Delete(id string) error
	Insert(user *types.User) error
}

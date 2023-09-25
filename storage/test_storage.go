package storage

import "github.com/elias-gill/poliapi/types"


type testStorage struct {
    table string
}

// returns a new connection to postgres
func NewTestStorage() *testStorage {
    return &testStorage{
        table: "usuarios",
    }
}

// get the user data of a given id
func (s *testStorage) GetById(user string) (*types.User, error) {
    return nil, nil
}

func (s *testStorage) Delete(user string) error {
    return nil
}

func (s *testStorage) Update(id string, user *types.User) error {
    return nil
}

func (s *testStorage) Insert(u *types.User) error {
    return nil
}

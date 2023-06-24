package store

import (
	"errors"
)

type PasswordData map[string]string

type Store struct {
	data PasswordData
}

func (store Store) ServiceExists(service string) bool {
	_, ok := store.data[service]
	return ok
}

func (store *Store) Write(service string, pwd string) {
	store.data[service] = pwd
}

func (store Store) Read(service string) (string, error) {
	pwd, ok := store.data[service]
	if !ok {
		return "", errors.New("service doesn't exist")
	} else {
		return pwd, nil
	}
}

func (store *Store) Delete(service string) {
	delete(store.data, service)
}

func (store Store) GetStore() PasswordData {
	return store.data
}

func NewPasswordStore(passwordData PasswordData) *Store {
	return &Store{data: passwordData}
}

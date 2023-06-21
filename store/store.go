package store

import (
	"encoding/json"
	"errors"

	"github.com/marcusprice/pw/util"
)

type Store struct {
	data map[string]string
}

func (store *Store) Init() {
	if util.DataFileExists() {
		file, err := util.ReadDataFile()
		if err != nil {
			panic(err)
		} else {
			json.Unmarshal(file, &store.data)
		}
	} else {
		util.CreateDataFile()
		store.data = make(map[string]string)
	}
}

func (store Store) ServiceExists(service string) bool {
	_, ok := store.data[service]
	return ok
}

func (store *Store) Add(service string, pwd string) {
	store.data[service] = pwd
}

func (store Store) Get(service string) (string, error) {
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

func (store Store) GetStore() map[string]string {
	return store.data
}

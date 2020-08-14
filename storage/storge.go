package storage

import (
	"encoding/gob"
	"os"
)

// Storage interface is used to store and retrieve arbitrary data to and from the storage backend
type Storage interface {
	Store(key string, value interface{}) error
	Load(key string, dest interface{}) error
}

type gobStorage struct {
	file *os.File
	data map[string]interface{}
}

// NewStorage returns a an implementation of the Storage interface
func NewStorage(storagePath string) (Storage, error) {

	f, err := os.OpenFile(storagePath, os.O_RDWR|os.O_CREATE, os.FileMode(0755))
	if err != nil {
		return nil, err
	}

	return &gobStorage{
		file: f,
		data: make(map[string]interface{}),
	}, nil
}
func (s *gobStorage) Store(key string, value interface{}) error {
	enc := gob.NewEncoder(s.file)

	s.data[key] = value

	if err := enc.Encode(s.data); err != nil {
		return err
	}

	return nil
}
func (s *gobStorage) Load(key string, dest interface{}) error {
	return nil
}

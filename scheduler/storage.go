package scheduler

import "github.com/pkg/errors"

type storage interface {
	Store(key string, data interface{}) error
	Load(key string, dest interface{}) error
}

var store storage

// SetStorage sets the storage to provided implementation of the storage interface
func SetStorage(s storage) {
	store = s
}

func persist(tasks []*Task) error {
	if store == nil {
		return nil
	}
	return store.Store("tasks", tasks)
}
func load(tasks *[]*Task) error {
	if store == nil {
		return errors.New("store is nil")
	}
	return store.Load("tasks", tasks)
}

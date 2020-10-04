package scheduler

import (
	"encoding/gob"
	"io"
	"os"
)

type tasks []*Task

var storageMedium io.ReadWriteCloser

func initStorage(options *Options) {
	gob.Register(new(Task))
	gob.Register(new(Schedule))
	gob.Register(new(Hour))
	gob.Register(new(Between))

	var err error
	storageMedium, err = os.OpenFile(options.StoragePath, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}

func store(tasks tasks) error {
	enc := gob.NewEncoder(storageMedium)

	return enc.Encode(tasks)
}

func load(tasks tasks) error {
	dec := gob.NewDecoder(storageMedium)

	return dec.Decode(tasks)
}

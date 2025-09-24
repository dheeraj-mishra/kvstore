package handlers

import "sync"

type datastore struct {
	dslock sync.RWMutex
	dsdata map[string]string
}

func newDataStore() *datastore {
	return &datastore{
		dslock: sync.RWMutex{},
		dsdata: make(map[string]string),
	}
}

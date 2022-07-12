package cache

import "nats/model"

type Cache struct {
	MP map[string]model.DataAll
}

// Init - инизиализация cache
func Init() *Cache {
	mp := new(Cache)
	mp.MP = make(map[string]model.DataAll)
	return mp
}

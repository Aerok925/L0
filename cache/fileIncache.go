package cache

import (
	"encoding/json"
	"errors"
	"nats/model"
	"sync"
)

func (cache *Cache) FileInCache(str []byte) error {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	var temp model.DataAll
	err := json.Unmarshal(str, &temp)
	if err != nil {
		return err
	}
	if temp.Order_uid == "" {
		return errors.New("Order_uid is missing")
	}
	if cache.MP[temp.Order_uid].Order_uid != "" {
		return errors.New("Такое уже есть!")
	}
	cache.MP[temp.Order_uid] = temp
	return nil
}

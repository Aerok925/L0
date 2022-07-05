package cache

import (
	"errors"
	"nats/model"
)

func (cache *Cache) FoundForUid(str string) (model.DataAll, error) {
	if cache.MP[str].Order_uid == "" {
		return cache.MP[str], errors.New("Not found")
	}
	return cache.MP[str], nil
}

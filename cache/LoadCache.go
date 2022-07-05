package cache

func (cache *Cache) LoadCache(strs *[]string) {
	for _, i := range *strs {
		cache.FileInCache([]byte(i))
	}
}

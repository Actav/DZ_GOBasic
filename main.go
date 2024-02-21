package main

import (
	"fmt"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{data: make(map[string]string)}
}

type cacheImpl struct {
	data map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	if v, ok := c.data[k]; ok {
		return v, ok
	}

	return "", false
}

func (c *cacheImpl) Set(k, v string) {
	c.data[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	if !ok {
		return "", false
	}
	d.cache.Set(k, v) // Сохраняем в кэш для будущего использования

	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)

	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))

	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}

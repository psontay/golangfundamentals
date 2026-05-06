package practice

type Cache[K comparable, V any] struct {
	items map[K]V
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{items: make(map[K]V)}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.items[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	v, ok := c.items[key]
	return v, ok
}

func (c *Cache[K, V]) Del(key K) {
	delete(c.items, key)
}

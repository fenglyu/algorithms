package cache

import (
	"sync"
)

type Cache struct {
	maxBytes int
}

type bucket struct {
	// read write lock
	mu sync.RWMutex

	// circle list
	chunks [][]byte

	// index dict
	m map[uint64]uint64

	// index
	idx uint64

	// The times that chunks got rewrritten
	gen uint64
}

func New(maxBytes int) *Cache {
	return &Cache{maxBytes: maxBytes}
}

func (c *Cache) Get(k []byte) []byte {

}

func (c *Cache) Set(k, v []byte) {

}

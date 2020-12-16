package util

import (
	"github.com/coocood/freecache"
	"runtime/debug"
)

const expire = 311040000

var cache *freecache.Cache

func init() {
	cache = freecache.NewCache(100 * 1024 * 1024)
	debug.SetGCPercent(20)
}

func Set(key, val []byte) error {
	return cache.Set(key, val, expire)
}

func Get(key []byte) (val []byte, err error) {
	return cache.Get(key)
}

func Del(key []byte) bool {
	return cache.Del(key)
}

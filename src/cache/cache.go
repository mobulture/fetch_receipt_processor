package cache

import (
	"fetch_receipt_processor/src/types"
)

type MemoryCache struct {
	data map[string]types.Receipt
}

var Cache *MemoryCache

func InitCache() {
	Cache = &MemoryCache{
		data: make(map[string]types.Receipt),
	}
}

func Set(key string, value types.Receipt) {
	Cache.data[key] = value
}

func Get(key string) (types.Receipt, bool) {
	val, exists := Cache.data[key]
	return val, exists
}

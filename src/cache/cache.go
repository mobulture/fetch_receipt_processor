package cache
import ("fetch_receipt_processor/src/types")


type MemoryCache struct {
    data map[string]types.Receipt
}

var Cache *MemoryCache

func initCache struct{
	Cache = &MemoryCache{
		data: make(map[string]types.Receipt),
	}
}

func Set(key, value string) {
    Cache.data[key] = value
}

func Get(key string) (string, bool) {
    val, exists := Cache.data[key]
    return val, exists
}


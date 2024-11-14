package types

type Item struct {
	ShortDescription string 
	Price            string 
}

type Receipt struct {
	Retailer    string  
	PurchaseDate string  
	PurchaseTime string  
	Items       []Item  
	Total       string  
}


type MemoryCache struct {
    data map[string]string
}


func newMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]string),
	}
}

func (cache *MemoryCache) Set(key, value string) {
    cache.data[key] = value
}

func (cache *MemoryCache) Get(key string) (string, bool) {
    val, exists := cache.data[key]
    return val, exists
}


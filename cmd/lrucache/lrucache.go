package lrucache

type CacheNode struct {
	key uint
	value uint
	prev *CacheNode
	next *CacheNode
}

type LRUCache struct {
	size uint
	hash   map[int]*CacheNode
	oldest *CacheNode
	newest *CacheNode
}

func NewLRUCache(capacity uint) *LRUCache {
	return &LRUCache{
		size: capacity,
		hash: map[int]*CacheNode{},
		oldest: nil,
		newest: nil,
	}
}

func (l *LRUCache) Set(k uint, v uint) {

	// condition where
	node := &CacheNode{
		key: k,
		value: v,
		prev: l.newest,
		next: nil,
	}
	l.newest = node
}

func (l *LRUCache) Get(k uint) int {
	return -1
}

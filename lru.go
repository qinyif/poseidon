package poseidon

import (
	"container/list"
	"sync"
)

type LRU struct {
	queue   *list.List
	storage *storage
	len     int64
	maxLen  int64
	clean   chan int
}

type queueItem struct {
	key string
}

type syncMapItem struct {
	ptr *list.Element
	val string
}

func NewLRU(maxLen int64) *LRU {
	lru := &LRU{
		queue:   list.New(),
		clean:   make(chan int, 100),
		storage: &storage{syncMap: new(sync.Map)},
		maxLen:  maxLen,
	}

	lru.initClean()

	return lru
}

func (lru *LRU) Get(key string) (string, bool) {
	if item, ok := lru.storage.Load(key); ok {
		lru.queue.MoveToFront(item.ptr)
		return item.val, true
	}

	return "", false
}

func (lru *LRU) Put(key string, value string) {
	item, ok := lru.storage.Load(key)
	if ok {
		lru.queue.MoveToFront(item.ptr)
	} else {
		elementPtr := lru.queue.PushFront(key)
		item = &syncMapItem{
			ptr: elementPtr,
			val: value,
		}
	}
	lru.storage.Store(key, item)

	incrLen := len(value)
	lru.clean <- incrLen
}

func (lru *LRU) Delete(key string) {
	item, ok := lru.storage.LoadAndDelete(key)
	if !ok {
		return
	}

	lru.queue.Remove(item.ptr)

	incrLen := -len(item.val)
	lru.clean <- incrLen
}

func (lru *LRU) initClean() {
	go func() {
		for {
			select {
			case incrLen := <-lru.clean:
				lru.len += int64(incrLen)
				if lru.len >= lru.maxLen && incrLen > 0 {
					back := lru.queue.Back()
					if back != nil {
						lru.queue.Remove(back)
						key := back.Value.(string)
						item, bo := lru.storage.Load(key)
						if bo {
							lru.maxLen -= int64(len(item.val))
							lru.storage.Delete(key)
						}
					}
				}
			}
		}
	}()
}

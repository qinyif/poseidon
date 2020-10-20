package poseidon

import "reflect"

type LocalStorage struct {
	sea *LRU
}

type Size int64

const (
	Bytes Size = 1
	KB         = 1000 * Bytes
	MB         = 1000 * KB
	GB         = 1000 * MB
)

func NewLocalStorage(size Size) *LocalStorage {
	maxLen := reflect.ValueOf(size).Int()

	return &LocalStorage{
		sea: NewLRU(maxLen),
	}
}

func (s *LocalStorage) SetItem(key string, value string) {
	s.sea.Put(key, value)
}

func (s *LocalStorage) GetItem(key string) (string, bool) {
	return s.sea.Get(key)
}

func (s *LocalStorage) RemoveItem(key string) {
	s.sea.Delete(key)
}

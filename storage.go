package poseidon

import (
	"sync"
)

type storage struct {
	syncMap *sync.Map
}

func (s *storage) Load(key string) (*syncMapItem, bool) {
	value, ok := s.syncMap.Load(key)
	if !ok {
		return nil, false
	}

	return value.(*syncMapItem), true
}

func (s *storage) Store(key string, item *syncMapItem) {
	s.syncMap.Store(key, item)
}

func (s *storage) LoadAndDelete(key string) (item *syncMapItem, loaded bool) {
	value, loaded := s.syncMap.LoadAndDelete(key)
	if !loaded {
		return nil, false
	}

	return value.(*syncMapItem), true
}

func (s *storage) Delete(key string) {
	s.syncMap.Delete(key)
}

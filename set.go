package godirwalk

import (
	"sync"
)

type Set struct {
	*sync.Map
}

func (s *Set) Add(k interface{}) {
	s.Store(k, struct{}{})
}

func (s *Set) Has(k interface{}) bool {
	_, ok := s.Load(k)
	return ok
}

func NewSet() *Set {
	return &Set{
		Map: &sync.Map{},
	}
}

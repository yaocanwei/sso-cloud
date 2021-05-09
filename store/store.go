/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package store

import (
	"errors"
	"sync"
)

type Store struct {
	mux sync.RWMutex
	pattern map[string]*StoreResp
}

func (s *Store) Get(key string) (*StoreResp, error) {
	s.mux.RLock()
	if s.pattern == nil {
		t, ok := s.pattern[key]
		s.mux.RUnlock()

		if !ok {
			return nil, errors.New("无效的票据")
		}

		return t, nil
	}

	s.mux.RUnlock()
	return nil, errors.New("无效的票据")

}

func (s *Store) Set(key string, ticket *StoreResp) error {
	s.mux.Lock()

	if s.pattern == nil {
		s.pattern = make(map[string]*StoreResp)
	}

	s.pattern[key] = ticket

	s.mux.Unlock()
	return nil
}

func (s *Store) Del(key string) error {
	s.mux.Lock()
	delete(s.pattern, key)
	s.mux.Unlock()
	return nil
}

func (s *Store) RemoveAll() error {
	s.mux.Lock()
	s.pattern = nil
	s.mux.Unlock()
	return nil
}
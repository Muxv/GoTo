package main

import "sync"

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func (s *URLStore) Get(key string) string {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	defer s.mu.Unlock()
	s.mu.Lock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
		mu:   sync.RWMutex{},
	}
}

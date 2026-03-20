package task

import (
	"sync"

	"github.com/google/uuid"
)

type Store struct {
	Tasks       []*Task
	mu          sync.RWMutex
	subscribers []chan Task
}

func NewStore() *Store { //returns initialized store
	new_store := &Store{}
	return new_store
}

func (s *Store) Add(t Task) {

	s.mu.Lock()
	defer s.mu.Unlock()

	t.ID = uuid.New() //assign unique id to task

	s.Tasks = append(s.Tasks, &t) // append it to internal slice

}

func (s *Store) List() []*Task {
	s.mu.RLock() //allows multiple people to view while still blocking if someone is adding one

	defer s.mu.RUnlock()

	return s.Tasks
}

func (s *Store) Subscribe() <-chan Task {

	s.mu.Lock()
	defer s.mu.Unlock()

	newChan := make(chan Task, 1)

	s.subscribers = append(s.subscribers, newChan)

	return newChan

}

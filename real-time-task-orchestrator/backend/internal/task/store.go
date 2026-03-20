package task

import "github.com/google/uuid"

type Store struct {
	Tasks []*Task
}

func NewStore() *Store { //returns initialized store
	new_store := &Store{}
	return new_store
}

func (s *Store) Add(t Task) {

	t.ID = uuid.New() //assign unique id to task

	s.Tasks = append(s.Tasks, &t) // append it to internal slice

}

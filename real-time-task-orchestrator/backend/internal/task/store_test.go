package task

import (
	"sync"
	"testing"
)

func TestStore_AddTask(t *testing.T) {
	s := Store{}

	ta := Task{Title: "Implement Subscriptions", Status: "Pending"}

	s.Add(ta)

	if len(s.Tasks) != 1 {
		t.Errorf("Expected length of store tasks = 1 got: %v", len(s.Tasks))
	}

	if s.Tasks[0].Title != "Implement Subscriptions" {
		t.Errorf("Wrong title: %v", s.Tasks[0].Title)
	}
}

func TestStore_ConcurrentAdd(t *testing.T) {
	s := Store{}

	wg := sync.WaitGroup{} //ensure all go routines finish before exiting the function

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done() //tells wg we are done
			ta := Task{Title: "Concurrent Task", Status: "Pending"}
			s.Add(ta)
		}()
	}

	wg.Wait() //wait for all routines to finish before proceeding

	if len(s.Tasks) != 1000 {
		t.Errorf("Expected 1000 tasks, got: %v", len(s.Tasks))

	}

}

func TestStore_Subscribe(t *testing.T) {

	s := Store{} // new store

	channel := s.Subscribe()


	go func(){
		s.Add(Task{Title: "Observer Task"})
	}


	case received := 

}

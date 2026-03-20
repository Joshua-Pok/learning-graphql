package gql

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Joshua-Pok/task-orchestrator/internal/gql"
	"github.com/Joshua-Pok/task-orchestrator/internal/gql/generated"
	"github.com/Joshua-Pok/task-orchestrator/internal/task"
)

func TestQuery_Tasks(t *testing.T) {
	s := &task.Store{}

	s.Add(task.Task{Title: "Fix Bug"})
	s.Add(task.Task{Title: "Write Docs"})

	resolver := &gql.Resolver{Store: s} //create the resolver

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: resolver}),
	)

	c := client.New(srv) //gql test Client

	var resp struct {
		Tasks []struct {
			Title string
		}
	}

	c.MustPost(`query { tasks { title } }`, &resp)

	if len(resp.Tasks) != 2 {
		t.Fatalf("expected 2 tasks got %d", len(resp.Tasks))
	}

	if resp.Tasks[0].Title != "Fix Bug" {
		t.Errorf("expected 'Fix Bug', got '%s'", resp.Tasks[0].Title)
	}
	if resp.Tasks[1].Title != "Write Docs" {
		t.Errorf("expected 'Write Docs', got '%s'", resp.Tasks[1].Title)
	}
}

func TestMutation_CreateTask(t *testing.T) {

	s := &task.Store{} //empty task.Store

	resolver := &gql.Resolver{Store: s}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: resolver}),
	)

	c := client.New(srv)

	var resp struct {
		Tasks []struct {
			Title string
		}
	}

	c.MustPost(`mutation {
  createTask(input: { title: "New Task" }) {
    id
    title
  }
}`, &resp)

	if s.List() == nil {
		t.Errorf("Task does not exist in memory")
	}

}

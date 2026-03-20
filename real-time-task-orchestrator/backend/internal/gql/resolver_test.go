package gql

import "github.com/Joshua-Pok/task-orchestrator/internal/task"

func TestQuery_Tasks() {

	s := task.Store{}

	queryString := "{ tasks { title } }"

	schema, err := gql.NewSchema

}

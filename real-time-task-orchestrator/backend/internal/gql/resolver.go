package gql

import (
	"github.com/Joshua-Pok/task-orchestrator/internal/task"
)

type Resolver struct {
	Store *task.Store
}

// type queryResolver struct{ *Resolver }
//
// func (r *Resolver) Query() generated.QueryResolver {
// 	return &queryResolver{r}
// }
//
// func (r *queryResolver) Tasks(ctx context.Context) ([]*generated.Task, error) {
// 	tasks := r.Store.List()
//
// 	var result []*generated.Task
//
// 	for _, t := range tasks {
// 		result = append(result, &generated.Task{
// 			ID:     t.ID.String(),
// 			Title:  t.Title,
// 			Status: t.Status,
// 		})
// 	}
//
// 	return result, nil
//
// }

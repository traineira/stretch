package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/traineira/stretch/cmd/graphql/graph/generated"
	"github.com/traineira/stretch/cmd/graphql/graph/model"
)

func (r *myMutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*model.Todo, error) {
	if todo.Done == nil {
		r.ToDo.Create(todo.Text, true)
	} else {
		r.ToDo.Create(todo.Text, false)
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *myMutationResolver) UpdateTodo(ctx context.Context, id string, updatedTodo model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// MyMutation returns generated.MyMutationResolver implementation.
func (r *Resolver) MyMutation() generated.MyMutationResolver { return &myMutationResolver{r} }

// MyQuery returns generated.MyQueryResolver implementation.
func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myMutationResolver struct{ *Resolver }
type myQueryResolver struct{ *Resolver }

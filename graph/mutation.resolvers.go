package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hayashi-ay/go-simple-graphql-api/graph/generated"
	"github.com/hayashi-ay/go-simple-graphql-api/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (bool, error) {
	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		User: &model.User{
			ID:   input.UserID,
			Name: "User " + input.UserID,
		},
	}

	r.todos = append(r.todos, todo)
	return true, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, todoID string, todoInput model.UpdateTodoInput) (bool, error) {
	for _, todo := range r.todos {
		if todo.ID == todoID {
			todo.Text = todoInput.Text
			todo.Done = todoInput.Done
			return true, nil
		}
	}
	return false, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

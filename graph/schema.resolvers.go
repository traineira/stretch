package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/traineira/stretch/graph/generated"
	"github.com/traineira/stretch/graph/model"
)

func (r *mutationResolver) CreateStory(ctx context.Context, input model.NewStory) (*model.Story, error) {
	story := &model.Story{
		Text:     input.Text,
		ID:       fmt.Sprintf("T%d", rand.Intn(30)),
		Category: "Test",
		UserID:   input.UserID,
	}
	r.stories = append(r.stories, story)
	return story, nil
}

func (r *queryResolver) Stories(ctx context.Context) ([]*model.Story, error) {
	return r.stories, nil
}

func (r *storyResolver) User(ctx context.Context, obj *model.Story) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Story returns generated.StoryResolver implementation.
func (r *Resolver) Story() generated.StoryResolver { return &storyResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type storyResolver struct{ *Resolver }

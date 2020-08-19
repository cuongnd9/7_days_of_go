package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/103cuong/uid"

	"github.com/103cuong/graphql/graph/generated"
	"github.com/103cuong/graphql/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	newLink := &model.Link{
		ID:      uid.Uid(16),
		Title:   input.Title,
		Address: input.Address,
		User: &model.User{
			ID:   uid.Uid(16),
			Name: "Cuong Duy Nguyen",
		},
	}
	return newLink, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	dummyLink := model.Link{
		ID:      uid.Uid(16),
		Title:   "Cuong Tran profile - GitHub",
		Address: "https://github.com/103cuong",
		User:    &model.User{ID: uid.Uid(16), Name: "Cuong Tran"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

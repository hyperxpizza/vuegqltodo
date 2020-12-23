package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hyperxpizza/vuegqltodo/server/database"
	"github.com/hyperxpizza/vuegqltodo/server/graph/generated"
	"github.com/hyperxpizza/vuegqltodo/server/graph/model"
	"github.com/hyperxpizza/vuegqltodo/server/helpers"
)

func (r *mutationResolver) CreateContact(ctx context.Context, input model.NewContact) (*model.Contact, error) {
	//validate input
	if helpers.ValidateContactInput(input) == false {
		return nil, fmt.Errorf("Error while walidating input")
	}

	//validate email
	err := helpers.ValidateEmail(input.Email)
	if err != nil {
		return nil, err
	}

	//insert
	contact, err := database.InsertContact(input)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *mutationResolver) UpdateContact(ctx context.Context, input model.UpdateContact) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllContacts(ctx context.Context) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

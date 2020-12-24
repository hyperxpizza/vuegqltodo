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
	contact, err := database.UpdateContact(input)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id int) (*model.DeleteResponse, error) {
	payload, err := database.DeleteContactByID(id)
	if err != nil {
		return nil, err
	}

	response := model.DeleteResponse{
		RowsAffected: *payload,
	}

	return &response, nil
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	contacts, err := database.GetAllContacts()
	if err != nil {
		return nil, err
	}

	return contacts, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

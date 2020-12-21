package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hyperxpizza/vuegqltodo/server/database"
	"github.com/hyperxpizza/vuegqltodo/server/graph/generated"
	"github.com/hyperxpizza/vuegqltodo/server/graph/model"
)

func (r *mutationResolver) CreateContact(ctx context.Context, input model.NewContact) (*model.Contact, error) {
	/* validate
	if helpers.ValidateContact(input) != true {
		return nil, fmt.Errorf("error while validating contact")
	}
	*/

	contact, err := database.InsertContact(input)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *mutationResolver) UpdateContact(ctx context.Context, input model.UpdateContact) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (int, error) {
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

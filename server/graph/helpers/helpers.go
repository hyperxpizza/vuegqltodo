package helpers

import "github.com/hyperxpizza/vuegqltodo/server/graph/model"

func ValidateContact(contact model.NewContact) bool {
	if len(contact.FirstName) > 100 || len(contact.LastName) > 100 || len(contact.Email) > 100 || len(contact.Phone) > 20 {
		return false
	}

	return true
}

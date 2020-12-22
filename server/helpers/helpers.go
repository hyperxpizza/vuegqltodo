package helpers

import (
	"log"
	"os"

	"github.com/badoux/checkmail"
	"github.com/hyperxpizza/vuegqltodo/server/graph/model"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func ValidateContactInput(contact model.NewContact) bool {
	if len(contact.FirstName) > 100 || len(contact.LastName) > 100 || len(contact.Email) > 100 || len(contact.Phone) > 20 {
		return false
	}

	return true
}

func ValidateEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	err = checkmail.ValidateHost(email)
	if err != nil {
		return err
	}

	return nil
}

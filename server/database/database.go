package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	//Postgresql driver
	"github.com/hyperxpizza/vuegqltodo/server/graph/model"
	_ "github.com/lib/pq"
)

var Database *sql.DB

const (
	host = "localhost"
	port = 5432
)

func InitDB(user, password, dbname string) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

	log.Println(psqlInfo)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	Database = database
	log.Println("Connected to the database")
}

func InsertContact(newContact model.NewContact) (*model.Contact, error) {
	stmt, err := Database.Prepare(`INSERT INTO contacts VALUES(DEFAULT, $1, $2, $3, $4, $5, $6)`)
	if err != nil {
		log.Print("1. error: ")
		log.Print(err)
		return nil, err
	}

	currentTime := time.Now()
	result, err := stmt.Exec(newContact.FirstName, newContact.LastName, newContact.Email, newContact.Phone, currentTime, currentTime)
	if err != nil {
		log.Print("2. error: ")
		log.Print(err)
		return nil, err
	}

	id2, err := result.RowsAffected()
	if err != nil {
		log.Print("3. error: ")
		log.Print(err)
		return nil, err
	}

	contact := model.Contact{
		ID:        string(id2),
		FirstName: newContact.FirstName,
		LastName:  newContact.LastName,
		Email:     newContact.Email,
		Phone:     newContact.Phone,
		CreatedAt: currentTime.String(),
		UpdatedAt: currentTime.String(),
	}

	return &contact, nil

}

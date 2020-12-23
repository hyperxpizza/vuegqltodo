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

var db *sql.DB

const (
	host = "localhost"
	port = 5432
)

func InitDB(user, password, dbname string) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println(psqlInfo)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = database
	log.Println("Connected to the database")

}

func InsertContact(input model.NewContact) (*model.Contact, error) {
	stmt, err := db.Prepare(`INSERT INTO contacts VALUES(DEFAULT, $1, $2, $3, $4, $5, $6) RETURNING id;`)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()

	var id int
	err = stmt.QueryRow(input.FirstName, input.LastName, input.Email, input.Phone, currentTime, currentTime).Scan(&id)
	if err != nil {
		return nil, err
	}

	contact := model.Contact{
		ID:        id,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	return &contact, nil

}

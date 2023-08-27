package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {

	host, _ := os.LookupEnv("DB_HOST")
	port, _ := os.LookupEnv("DB_PORT")
	user, _ := os.LookupEnv("DB_USER")
	password, _ := os.LookupEnv("DB_PASSWORD")
	dbname, _ := os.LookupEnv("DB_NAME")

	connStr := "host=" + host + " port=" +
		port + " user=" + user + " password=" +
		password + " dbname=" + dbname + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}

func Close(db *sql.DB) {
	log.Println("Closing database")
	db.Close()
}

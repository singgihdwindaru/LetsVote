package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

func SetupDB() *sql.DB {
	var host string = os.Getenv("DB_HOST")
	var dbName string = os.Getenv("DB_NAME")
	var user string = os.Getenv("DB_USERNAME")
	var password string = os.Getenv("DB_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		// ... handle error
		panic(err)
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(0)
	return db
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if err != nil {
			log.Panic(errorRollback)
		}
		panic(err)
	} else {
		errorCommit := tx.Commit()
		if err != nil {
			log.Panic(errorCommit)
		}
	}
}

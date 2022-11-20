package main

import (
	"database/sql"
	"log"

	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/config"
)

var (
	DB *sql.DB
)

func main() {
	DB = config.SetupDB()
	defer DB.Close()

	r := config.SetGinRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

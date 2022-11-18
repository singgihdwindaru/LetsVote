package main

import (
	"fmt"
	"log"

	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/config"
)

func main() {
	fmt.Println("Secure Vote")
	r := config.SetGinRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

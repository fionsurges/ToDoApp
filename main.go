package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fionwan/todoApp/database"
	"github.com/fionwan/todoApp/handlers"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080", nil
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	database.InitDB()
	router := handlers.NewRouter()
	
	fmt.Println("server listening on port", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

package main 

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	InitDB() 
	router := NewRouter()

	fmt.Println("server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}


package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := "Hello les ziGOtos ;)...... :-|"
	fmt.Println(message)
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Main launch")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	appEnv := os.Getenv("APP_ENV")
	message := "Hello World"

	log.Println("APP_ENV:", appEnv)

	if appEnv == "prod" {
		message = "Hello World from production"
	}

	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
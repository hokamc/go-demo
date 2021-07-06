package main

import (
	di "example.com/dependecy-injection"
	"log"
	"net/http"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Elodie")

	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(di.MyGreetHandler)))
}

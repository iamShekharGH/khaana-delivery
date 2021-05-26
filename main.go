package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hi there.")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on \t:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	initialInit()
}

func initialInit() {
	http.HandleFunc("/", indexFunc)
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	s := `
		<title>THis is ShekharGH</title>
		<h1>Init is loading</h1>
	`
	io.WriteString(w, s)
}

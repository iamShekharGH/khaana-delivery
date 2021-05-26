package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template
var env string

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env = "https://khaana-mangva-do.herokuapp.com/"

	fmt.Println("Hi there.")
	log.Println("Listening on \t:", port)

	initialInit()

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func initialInit() {
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/test/", test)
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	s := `
		<html>
		<title>THis is ShekharGH</title>
		
		<body>
		<h1>Init is loading</h1>
		</body>
		</html>
	`

	io.WriteString(w, s)
}

func test(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "This really shud work!!")
}

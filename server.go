package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type page struct {
	Title   string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := page{
		Title:   "Go hs",
		Message: "Hey, Gopher!",
	}

	t := template.Must(template.ParseFiles("./index.tmpl"))

	err := t.ExecuteTemplate(w, "index.tmpl", p)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening")
	http.ListenAndServe(":8000", nil)
}

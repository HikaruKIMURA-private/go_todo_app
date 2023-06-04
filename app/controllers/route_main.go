package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "hello world")
}
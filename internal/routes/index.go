package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Println("GET Index")
	tmpl:= template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

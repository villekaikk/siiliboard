package api

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GET /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Println("GET /")
	tmpl := template.Must(template.ParseFiles("templates/pages/page_index.html"))
	tmpl.Execute(w, nil)
}

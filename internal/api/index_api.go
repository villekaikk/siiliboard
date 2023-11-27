package api

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GET /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tmpl := template.Must(template.ParseFiles("templates/pages/page_index.html"))
	tmpl.Execute(w, nil)
}

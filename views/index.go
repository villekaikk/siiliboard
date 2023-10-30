package views

import (
	"text/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tmpl:= template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

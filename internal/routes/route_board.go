package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"

	"siiliboard/internal/database"
	"siiliboard/internal/marshal"
)

// GET /boards
func GetBoardsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Println("GET Boards")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// GET /board/<id>
func GetBoardHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	log.Printf("GET Board %s\n", id)
	fmt.Fprintf(w, "Requested board %s\n", id)
}

// POST /boards
func CreateBoardHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	log.Printf("POST Board\n")
	b := &marshal.BoardRequest{}
	err := readBodyToModel(r, b)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, err.Error())
		return
	}

	id, err := database.CreateBoard(b)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - Failed to create a new board: %s", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to create a new board")
		return
	}

	status := http.StatusCreated
	log.Printf("%d - New board created: %d - %s", status, id, b.Name)
	w.WriteHeader(status)
}

// DELETE /boards
func RemoveAllBoardsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("DELETE Boards\n")
}

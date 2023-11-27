package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"

	"siiliboard/internal/database"
	"siiliboard/internal/marshal"
)

// GET /boards
func GetBoardsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	boards, err := database.GetBoards()

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch boards")
		return
	}

	status := http.StatusOK
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/fragments/boards.html"))
	for _, v := range boards {
		tmpl.Execute(w, v)
	}
}

// GET /boards/:bid
func GetBoardHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("bid")
	idd, err := strconv.Atoi(id)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid board id %s", id)
		return
	}

	b, err := database.GetBoard(idd)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch boards")
		return
	}

	status := http.StatusOK
	w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles("templates/pages/page_board.html"))
	tmpl.Execute(w, b)
}

// POST /boards
func CreateBoardHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	err := r.ParseForm()

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid form data")
		return
	}

	b, err := marshal.NewBoardRequest(r, *decoder)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Unable to deserialize form")
		return
	}

	err = b.Validate()

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, err.Error())
		return
	}

	board, err := database.CreateBoard(b)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to create a new board")
		return
	}

	status := http.StatusCreated
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/fragments/boards.html"))
	tmpl.Execute(w, board)
}

// DELETE /boards/:bid
func RemoveAllBoardsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

// GET /newboard
func GetNewBoardModal(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	status := http.StatusOK
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/modals/new_board_modal.html"))
	tmpl.Execute(w, nil)
}

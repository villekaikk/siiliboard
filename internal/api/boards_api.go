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

	endpoint := fmt.Sprintf("GET %s", r.URL)
	log.Println(endpoint)
	boards, err := database.GetBoards()

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch boards")
		return
	}

	status := http.StatusOK
	log.Printf("%d - %s {%d}\n", status, endpoint, len(boards))
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/fragments/boards.html"))
	for _, v := range boards {
		tmpl.Execute(w, v)
	}
}

// GET /boards/:bid
func GetBoardHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("bid")
	endpoint := fmt.Sprintf("GET %s", r.URL)
	log.Println(endpoint)

	idd, err := strconv.Atoi(id)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid board id %s", id)
		return
	}

	b, err := database.GetBoard(idd)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch boards")
		return
	}

	status := http.StatusOK
	log.Printf("%d - %s", status, endpoint)
	w.WriteHeader(status)
	//fmt.Fprintf(w, "Boards: %v", b)
	tmpl := template.Must(template.ParseFiles("templates/pages/page_board.html"))
	tmpl.Execute(w, b)
}

// POST /boards
func CreateBoardHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	endpoint := fmt.Sprintf("POST %s", r.URL)
	log.Println(endpoint)

	err := r.ParseForm()

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid form data")
		return
	}

	b, err := marshal.NewBoardRequest(r, *decoder)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Unable to deserialize form")
		return
	}

	err = b.Validate()

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, err.Error())
		return
	}

	board, err := database.CreateBoard(b)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to create a new board")
		return
	}

	status := http.StatusCreated
	log.Printf("%d - %s - %s\n", status, endpoint, board.Name)
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/fragments/boards.html"))
	tmpl.Execute(w, board)
}

// DELETE /boards/:bid
func RemoveAllBoardsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("DELETE %s\n", r.URL)
}

// GET /newboard
func GetNewBoardModal(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	endpoint := fmt.Sprintf("GET %s", r.URL)
	log.Println(endpoint)

	status := http.StatusOK
	log.Printf("%d - %s\n", status, endpoint)
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/modals/new_board_modal.html"))
	tmpl.Execute(w, nil)
}

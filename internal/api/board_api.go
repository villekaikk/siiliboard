package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"siiliboard/internal/database"
	"siiliboard/internal/marshal"
)

// GET /boards
func GetBoardsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Println("GET Boards")
	boards, err := database.GetBoards()

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - Failed to fetch boards", errCode)
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch boards")
		return
	}

	status := http.StatusOK
	log.Printf("%d - Get all boards (%d)", status, len(boards))
	w.WriteHeader(status)
	fmt.Fprintf(w, "Boards: %v", boards)
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

	board, err := database.CreateBoard(b)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - Failed to create a new board: %s", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to create a new board")
		return
	}

	status := http.StatusCreated
	log.Printf("%d - New board created: %d - %s", status, board.Id, board.Name)
	w.WriteHeader(status)
	fmt.Fprintf(w, "Board: %v", board)
}

// DELETE /boards
func RemoveAllBoardsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("DELETE Boards\n")
}

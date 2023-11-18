package api

import (
	"fmt"
	"log"
	"net/http"
	"siiliboard/internal/database"
	"siiliboard/internal/marshal"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// GET /boards/:bid/tickets|?state=["todo", "inprogress", "done"]
func GetTicketsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	board_id := params.ByName("bid")
	state := r.URL.Query().Get("state")
	endpoint := fmt.Sprintf("GET %s", r.URL)
	log.Println(endpoint)

	idd, err := strconv.Atoi(board_id)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid board id %s", board_id)
		return
	}

	tickets, err := database.GetTickets(idd, state)

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

	tmpl := template.Must(template.ParseFiles("templates/fragments/ticket.html"))
	for _, v := range tickets {
		tmpl.Execute(w, v)
	}
}

// GET /boards/:bid/ticket/:tid
func GetTicketHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	board_id := params.ByName("bid")
	ticket_id := params.ByName("tid")
	endpoint := fmt.Sprintf("GET %s", r.URL)
	log.Println(endpoint)

	bid, err := strconv.Atoi(board_id)

	if err != nil || bid < 1 {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid board id %s\n", board_id)
		return
	}

	tid, err := strconv.Atoi(ticket_id)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid ticket id %s", board_id)
		return
	}

	t, err := database.GetTicket(bid, tid)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch ticket %d", tid)
		return
	}

	status := http.StatusOK
	log.Printf("%d - %s", status, endpoint)
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/fragments/ticket.html"))
	tmpl.Execute(w, t)
}

// POST /boards/:bid/tickets
func CreateTicketHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	endpoint := fmt.Sprintf("POST %s", r.URL)

	board_id := params.ByName("bid")
	bid, err := strconv.Atoi(board_id)
	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid board id %s\n", board_id)
		return
	}

	log.Println(endpoint)

	t := &marshal.TicketRequest{Board: bid}
	err = readBodyToModel(r, t)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, err.Error())
		return
	}

	ticket, err := database.CreateTicket(t)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to create a new ticket")
		return
	}

	status := http.StatusCreated
	log.Printf("%d - %s - %d\n", status, endpoint, ticket.Id)
	w.WriteHeader(status)
	fmt.Fprintf(w, "Board: %v\n", ticket)
}

// GET /boards/:bid/newticket
func GetNewTicketModal(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	board_id := params.ByName("bid")
	endpoint := fmt.Sprintf("GET %s", r.URL)
	log.Println(endpoint)

	bid, err := strconv.Atoi(board_id)

	if err != nil || bid < 1 {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid board id %s", board_id)
		return
	}

	b, err := database.GetBoard(bid)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch boards\n")
		return
	}

	status := http.StatusOK
	log.Printf("%d - %s", status, endpoint)
	w.WriteHeader(status)

	tmpl := template.Must(template.ParseFiles("templates/modals/new_ticket_modal.html"))
	tmpl.Execute(w, b)
}

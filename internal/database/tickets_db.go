package database

import (
	"fmt"
	"log"

	"siiliboard/internal/domain"
	"siiliboard/internal/marshal"
	"siiliboard/internal/utils"
)

func GetTickets(board_id int, state string) ([]domain.Ticket, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	tickets := []domain.Ticket{}

	query := fmt.Sprintf("SELECT * FROM ticket WHERE board_id=%d", board_id)

	if !utils.IsStringEmpty(state) {
		query = fmt.Sprintf("%s AND state='%s'", query, state)
	}

	err = db.Database.Select(&tickets, query)

	if err != nil {
		log.Printf("Error querying tickets of board %d from the database: %s\n", board_id, err.Error())
		return nil, err
	}

	return tickets, nil
}

func GetTicket(board_id int, ticket_id int) (*domain.Ticket, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	ticket := domain.NewTicket()

	q := "SELECT * FROM ticket WHERE ticket_id = $1 AND board_id = $2"
	err = db.Database.Get(ticket, q, ticket_id, board_id)

	if err != nil {
		log.Printf("Error querying ticket %d from the database: %s\n", ticket_id, err.Error())
		return nil, err
	}

	ticket.Assignee = nil
	ass, err := GetUser(ticket.AssigneeId)
	if err != nil {
		log.Printf("Error querying ticket's %d assignee %d from the database: %s\n", ticket_id, ticket.AssigneeId, err.Error())
	} else {
		ticket.Assignee = ass
	}

	ticket.Author = nil
	auth, err := GetUser(ticket.AuthorId)
	if err != nil {
		log.Printf("Error querying ticket's %d author %d from the database: %s\n", ticket_id, ticket.AuthorId, err.Error())
	} else {
		ticket.Author = auth
	}

	return ticket, nil
}

func CreateTicket(tr *marshal.TicketRequest) (*domain.Ticket, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	var ticket_id int
	query := `INSERT INTO ticket (name, description, author, assignee, board_id) VALUES ($1, $2, $3, $4, $5) RETURNING ticket_id`
	err = db.Database.QueryRow(query, tr.Name, tr.Description, tr.Author, tr.Assignee, tr.Board).Scan(&ticket_id)

	if err != nil {
		log.Printf("Unable to create new ticket: %s\n", err.Error())
		return nil, err
	}

	t, err := GetTicket(tr.Board, ticket_id)

	if err != nil {
		log.Printf("Unable to query newly made ticket: %s\n", err.Error())
		return nil, err
	}

	return t, nil
}

package marshal

import (
	"fmt"
	"siiliboard/internal/utils"
)

type TicketRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      int    `json:"author"`
	Assignee    int    `json:"assignee"`
	Board       int
}

func (t *TicketRequest) Validate() error {

	if utils.IsStringEmpty(t.Name) {
		return fmt.Errorf("Could not serialize ticket 'name'")
	}

	if utils.IsStringEmpty(t.Description) {
		return fmt.Errorf("Could not serialize ticket 'description'")
	}

	if t.Author < -1 {
		return fmt.Errorf("Invalid author id, can't be <1")
	}

	return nil
}

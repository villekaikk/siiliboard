package marshal

import (
	"fmt"
	"log"
	"net/http"
	"siiliboard/internal/utils"

	"github.com/gorilla/schema"
)

type BoardRequest struct {
	Name string `schema:"board-name"`
}

func (b BoardRequest) Validate() error {

	if utils.IsStringEmpty(b.Name) {
		return fmt.Errorf("Could not serialize ticket 'name'")
	}

	return nil
}

func NewBoardRequest(r *http.Request, d schema.Decoder) (*BoardRequest, error) {

	var b BoardRequest
	err := d.Decode(&b, r.PostForm)

	if err != nil {
		log.Printf("Error decoding BoardRequest - %s", err.Error())
		return nil, err
	}

	return &b, nil
}

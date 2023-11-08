package marshal

import (
	"errors"
	"siiliboard/internal/utils"
)

type BoardRequest struct {
	Name string `json:"name"`
}

func (b BoardRequest) Validate() error {

	if utils.IsEmpty(b.Name) {
		return errors.New("Could not serialize 'name'")
	}

	return nil
}

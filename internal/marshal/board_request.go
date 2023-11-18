package marshal

import (
	"fmt"
	"siiliboard/internal/utils"
)

type BoardRequest struct {
	Name string `json:"name"`
}

func (b BoardRequest) Validate() error {

	if utils.IsStringEmpty(b.Name) {
		return fmt.Errorf("Could not serialize ticket 'name'")
	}

	return nil
}

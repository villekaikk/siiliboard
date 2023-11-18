package marshal

import (
	"fmt"
	"siiliboard/internal/utils"
)

type UserRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func (u *UserRequest) Validate() error {

	if utils.IsStringEmpty(u.Name) {
		return fmt.Errorf("Could not serialize user 'name'")
	}

	if utils.IsStringEmpty(u.DisplayName) {
		return fmt.Errorf("Could not serialize user 'display_name'")
	}

	return nil
}

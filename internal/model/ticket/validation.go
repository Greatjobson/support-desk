package ticket

import (
	"errors"
	"strings"
)

func ValidateTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("title is empty")
	}

	if len([]rune(title)) > 100 {
		return errors.New("title is too long")
	}
	return nil
}

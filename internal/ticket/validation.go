package ticket

import (
	"errors"
	"strings"
)

func ValidateTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("заголовок не должен быть пустым")
	}

	if len([]rune(title)) > 100 {
		return errors.New("заголовок не должен быть больше 100 символов")
	}
	return nil
}

package message

import (
	"errors"
	"strings"
)

func ValidateMessageBody(body string) error {
	body = strings.TrimSpace(body)
	if body == "" {
		return errors.New("body is empty")
	}
	if len([]rune(body)) > 4099 { //думаю хватит + надо предупреждать в фронтенде что осталось 100 символов и тп
		return errors.New("body text is too long")
	}
	return nil
}

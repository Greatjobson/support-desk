package message

import (
	"strings"
	"testing"
)

func TestValidateMessageBody(t *testing.T) {

	t.Run("empty message", func(t *testing.T) {
		err := ValidateMessageBody("  ")
		if err == nil {
			t.Error("Expected error of empty message, got nil")
		}
	})
	t.Run("long message + 1", func(t *testing.T) {
		err := ValidateMessageBody(strings.Repeat("A", 4100))
		if err == nil {
			t.Error("Expected error, got nil, check placeholder of max length")
		}
	})
	t.Run("placeholder fitt", func(t *testing.T) {
		err := ValidateMessageBody(strings.Repeat("A", 4099))
		if err != nil {
			t.Error("Expected nil , got error, check placeholder of max length")
		}
	})
	t.Run("example 1", func(t *testing.T) {
		err := ValidateMessageBody("I got problem with my very large penis and blue balls")
		if err != nil {
			t.Error("Expected nil , got error, check placeholder of max length")
		}
	})
}

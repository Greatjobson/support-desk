package main

import (
	"fmt"
	"strings"

	"github.com/Greatjobson/support-desk/internal/ticket"
)

func main() {
	ticket1 := ticket.Ticket{
		ID:          1,
		Title:       strings.Repeat("Я", 10),
		Description: "her otvalilsya,help",
		Status:      "open",
	}

	err := ticket.ValidateTitle(ticket1.Title)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ticket1)
}

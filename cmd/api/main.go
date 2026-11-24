package main

import (
	"fmt"
	"strings"

	ticket2 "github.com/Greatjobson/support-desk/internal/model/ticket"
)

func main() {
	ticket1 := ticket2.Ticket{
		ID:          "1",
		Title:       strings.Repeat("Я", 10),
		Description: "her otvalilsya,help",
		Status:      "open",
	}

	err := ticket2.ValidateTitle(ticket1.Title)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ticket1)
}

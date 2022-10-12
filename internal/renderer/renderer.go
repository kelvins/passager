package renderer

import (
	"fmt"
	"io"
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/olekukonko/tablewriter"
)

// PrintCredentials render a table with one or more credentials to the provided io.Writer.
func PrintCredentials(w io.Writer, data interface{}) {
	var credentials []models.Credential
	switch value := data.(type) {
	case models.Credential:
		credentials = append(credentials, value)
	case []models.Credential:
		credentials = append(credentials, value...)
	default:
		log.Fatal("Invalid value to print")
	}

	if len(credentials) == 0 {
		fmt.Println("No credentials found!")
		return
	}

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Name", "Login", "Password", "Description"})

	for _, v := range credentials {
		table.Append([]string{v.Name, v.Login, v.Password, v.Description})
	}
	table.Render()
}

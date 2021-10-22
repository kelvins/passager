package renderer

import (
	"fmt"
	"log"
	"os"

	"github.com/kelvins/passager/internal/models"
	"github.com/olekukonko/tablewriter"
)

func PrintCredentials(data interface{}) {
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Login", "Password", "Description"})

	for _, v := range credentials {
		table.Append([]string{v.Name, v.Login, v.Password, v.Description})
	}
	table.Render()
}

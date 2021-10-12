package renderer

import (
	"os"
	"log"

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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Login", "Password"})

	for _, v := range credentials {
		table.Append([]string{v.Name, v.Login, v.Password})
	}
	table.Render()
}

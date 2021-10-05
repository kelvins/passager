package cli

import (
	"fmt"

	"github.com/kelvins/passager/internal/db"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all database credentials",
	Long:  "List all database credentials",
	Run:   listCredentials,
}

func listCredentials(cmd *cobra.Command, args []string) {
	credentials, err := db.ReadAll()
	if err == nil {
		for _, credential := range credentials {
			fmt.Printf("Login: %s\nPassword: %s\n\n", credential.Login, credential.Password)
		}
	} else {
		fmt.Println("Not found")
	}
}

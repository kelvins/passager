package commands

import (
	"fmt"

	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
)

func ListCmdFactory() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all database credentials",
		Long:  "List all database credentials",
		Run:   listCmdRun,
	}
	return listCmd
}

func listCmdRun(cmd *cobra.Command, args []string) {
	credentials, err := models.ReadAll()
	if err == nil {
		fmt.Printf("| %-24s| %-24s| %-24s|\n", "Name", "Login", "Password")
		for _, credential := range credentials {
			fmt.Println(credential.String())
		}
	} else {
		fmt.Println("Not found")
	}
}

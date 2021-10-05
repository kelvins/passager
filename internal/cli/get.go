package cli

import (
	"fmt"

	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
)

func GetCmdFactory() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Set credentials to the database",
		Long:  "Set credentials to the database",
		Args:  cobra.MinimumNArgs(1),
		Run:   getCmdRun,
	}
	return getCmd
}

func getCmdRun(cmd *cobra.Command, args []string) {
	credential, err := models.Read(args[0])
	if err == nil {
		fmt.Printf("Login: %s\nPassword: %s\n", credential.Login, credential.Password)
	} else {
		fmt.Println("Not found")
	}
}

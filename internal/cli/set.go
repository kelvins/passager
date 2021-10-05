package cli

import (
	"fmt"
	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
)

func SetCmdFactory() *cobra.Command {
	var setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set credentials to the database",
		Long:  "Set credentials to the database",
		Args:  cobra.MinimumNArgs(3),
		Run:   setCmdRun,
	}
	return setCmd
}

func setCmdRun(cmd *cobra.Command, args []string) {
	credential := models.Credential{
		Name:     args[0],
		Login:    args[1],
		Password: args[2],
	}
	err := models.Create(&credential)
	if err == nil {
		fmt.Println("Created")
	} else {
		fmt.Println("Error creating the credential")
	}
}

package commands

import (
	"fmt"
	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
)

func DeleteCmdFactory() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete credentials from the database",
		Long:  "Delete credentials from the database",
		Args:  cobra.MinimumNArgs(1),
		Run:   deleteCmdRun,
	}
	return deleteCmd
}

func deleteCmdRun(cmd *cobra.Command, args []string) {
	err := models.Delete(args[0])
	if err == nil {
		fmt.Println("Deleted")
	} else {
		fmt.Println("Error deleting the credential")
	}
}

package commands

import (
	"fmt"
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
)

// DeleteCmdFactory is responsible for creating the passager delete sub-command.
// This command is responsible for deleting existing credentials from the database.
func DeleteCmdFactory() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete [NAME]",
		Short: "Delete a credential from the database",
		Args:  cobra.ExactArgs(1),
		Run:   deleteCmdRun,
	}
	return deleteCmd
}

func deleteCmdRun(cmd *cobra.Command, args []string) {
	name := args[0]
	err := models.Delete(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Credential %s successfully deleted!\n", name)
}

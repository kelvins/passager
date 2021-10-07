package commands

import (
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

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
	color.Green("Credential %s was successfully deleted!", name)
}

package commands

import (
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

func SetCmdFactory() *cobra.Command {
	var setCmd = &cobra.Command{
		Use:   "set [NAME] [LOGIN] [PASSWORD]",
		Short: "Set credentials to the database",
		Args:  cobra.ExactArgs(3),
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
	if err := models.Create(&credential); err != nil {
		log.Fatal(err)
	}
	color.Green("Credential %s successfully created!", args[0])
}

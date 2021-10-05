package commands

import (
	"fmt"
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/spf13/cobra"
)

func GetCmdFactory() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get [NAME]",
		Short: "Set credentials to the database",
		Args:  cobra.ExactArgs(1),
		Run:   getCmdRun,
	}
	return getCmd
}

func getCmdRun(cmd *cobra.Command, args []string) {
	credentials, err := models.Read(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("| %-24s| %-24s| %-24s|\n", "Name", "Login", "Password")
	for _, credential := range credentials {
		fmt.Println(credential.String())
	}
}

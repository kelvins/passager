package commands

import (
	"log"

	"github.com/fatih/color"
	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/spf13/cobra"
)

func AddCmdFactory() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add [NAME] [LOGIN] [PASSWORD]",
		Short: "Add credentials to the database",
		Args:  cobra.ExactArgs(3),
		Run:   addCmdRun,
	}

	addCmd.Flags().StringP("key", "k", "", "Key to encrypt/decrypt data using AES")

	return addCmd
}

func addCmdRun(cmd *cobra.Command, args []string) {
	password := args[2]
	key, err := cmd.Flags().GetString("key")
	if len(key) > 0 {
		if len(key) != 16 || err != nil {
			log.Fatal("Invalid key")
		} else {
			password = crypto.Encrypt(key, password)
		}
	}
	credential := models.Credential{
		Name:     args[0],
		Login:    args[1],
		Password: password,
	}
	if err := models.Create(&credential); err != nil {
		log.Fatal(err)
	}
	color.Green("Credential %s successfully created!", args[0])
}
package commands

import (
	"fmt"
	"log"

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

	addCmd.Flags().StringP("key", "k", crypto.EncryptionKey(), "Key to encrypt/decrypt data")
	addCmd.Flags().StringP("description", "d", "", "Description related to the credentials")

	return addCmd
}

func addCmdRun(cmd *cobra.Command, args []string) {
	key, err := cmd.Flags().GetString("key")
	if err != nil {
		log.Fatal("Invalid key")
	}
	password := crypto.Encrypt(args[2], key)
	description, _ := cmd.Flags().GetString("description")
	credential := models.Credential{
		Name:        args[0],
		Login:       args[1],
		Password:    password,
		Description: description,
	}
	if err := models.Create(&credential); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Credential %s successfully created!", args[0])
}

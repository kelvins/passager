package commands

import (
	"fmt"
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/spf13/cobra"
)

// AddCmdFactory is responsible for creating the passager add sub-command.
// This command is responsible for adding new credentials to the database.
func AddCmdFactory() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add [NAME]",
		Short: "Add credentials to the database",
		Args:  cobra.ExactArgs(1),
		Run:   addCmdRun,
	}
	addCmd.Flags().StringP("login", "l", "", "Credential login")
	addCmd.Flags().StringP("password", "p", "", "Credential password (required)")
	addCmd.Flags().StringP("key", "k", crypto.EncryptionKey(), "Key to encrypt/decrypt data")
	addCmd.Flags().StringP("description", "d", "", "Description related to the credentials")
	addCmd.MarkFlagRequired("password")
	return addCmd
}

func getStringFlag(cmd *cobra.Command, flag string) string {
	value, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func addCmdRun(cmd *cobra.Command, args []string) {
	key := getStringFlag(cmd, "key")
	password := getStringFlag(cmd, "password")
	credential := models.Credential{
		Name:        args[0],
		Login:       getStringFlag(cmd, "login"),
		Password:    crypto.Encrypt(password, key),
		Description: getStringFlag(cmd, "description"),
	}
	if err := models.Create(&credential); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Credential %s successfully created!\n", args[0])
}

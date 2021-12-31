package commands

import (
	"fmt"
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/spf13/cobra"
)

func EditCmdFactory() *cobra.Command {
	var editCmd = &cobra.Command{
		Use:   "edit [NAME]",
		Short: "Edit existing credentials",
		Args:  cobra.ExactArgs(1),
		Run:   editCmdRun,
	}
	editCmd.Flags().StringP("login", "l", "", "Credential login")
	editCmd.Flags().StringP("password", "p", "", "Credential password")
	editCmd.Flags().StringP("description", "d", "", "Credential description")
	editCmd.Flags().StringP("key", "k", crypto.EncryptionKey(), "Key to encrypt/decrypt data")
	return editCmd
}

func getStringFlagOrDefault(cmd *cobra.Command, flag, defaultValue string) string {
	value, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatal(err)
	}
	if value != "" {
		return value
	}
	return defaultValue
}

func editCmdRun(cmd *cobra.Command, args []string) {
	credential, err := models.ReadFirst(args[0])
	if err != nil {
		log.Fatal(err)
	}
	credential.Login = getStringFlagOrDefault(cmd, "login", credential.Login)
	credential.Password = getStringFlagOrDefault(cmd, "password", credential.Password)
	credential.Description = getStringFlagOrDefault(cmd, "description", credential.Description)
	// TODO: save credential to database
	fmt.Fprintf(cmd.OutOrStdout(), "Credential %s successfully updated!\n", args[0])
}

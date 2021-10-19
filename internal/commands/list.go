package commands

import (
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/internal/renderer"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/spf13/cobra"
)

func ListCmdFactory() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all database credentials",
		Run:   listCmdRun,
	}

	listCmd.Flags().StringP("key", "k", "#encryption-key#", "Key to encrypt/decrypt data")

	return listCmd
}

func listCmdRun(cmd *cobra.Command, args []string) {
	key, err := cmd.Flags().GetString("key")
	if err != nil {
		log.Fatal("Invalid key")
	}
	credentials, err := models.ReadAll("")
	if err != nil {
		log.Fatal(err)
	}
	var newCredentials []models.Credential
	for _, credential := range credentials {
		credential.Password = crypto.Decrypt(credential.Password, key)
		newCredentials = append(newCredentials, credential)
	}
	renderer.PrintCredentials(newCredentials)
}

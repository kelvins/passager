package commands

import (
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/internal/renderer"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/spf13/cobra"
)

func GetCmdFactory() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get [NAME]",
		Short: "Get credentials from the database",
		Args:  cobra.ExactArgs(1),
		Run:   getCmdRun,
	}

	getCmd.Flags().StringP("key", "k", "#encryption-key#", "Key to encrypt/decrypt data")

	return getCmd
}

func getCmdRun(cmd *cobra.Command, args []string) {
	credentials, err := models.ReadAll(args[0])
	if err != nil {
		log.Fatal(err)
	}
	key, err := cmd.Flags().GetString("key")
	if err != nil {
		log.Fatal("Invalid key")
	}
	var newCredentials []models.Credential
	for _, credential := range credentials {
		credential.Password = crypto.Decrypt(credential.Password, key)
		newCredentials = append(newCredentials, credential)
	}
	renderer.PrintCredentials(newCredentials)
}

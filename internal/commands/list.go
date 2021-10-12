package commands

import (
	"fmt"
	"log"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/spf13/cobra"
)

func ListCmdFactory() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all database credentials",
		Run:   listCmdRun,
	}

	listCmd.Flags().StringP("key", "k", "", "Key to encrypt/decrypt data")

	return listCmd
}

func listCmdRun(cmd *cobra.Command, args []string) {
	key, err := cmd.Flags().GetString("key")
	if err != nil {
		log.Fatal("Invalid key")
	}
	credentials, err := models.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("| %-24s| %-24s| %-24s|\n", "Name", "Login", "Password")
	for _, credential := range credentials {
		if len(key) > 0 {
			credential.Password = crypto.Decrypt(credential.Password, key)
		}
		fmt.Println(credential.String())
	}
}

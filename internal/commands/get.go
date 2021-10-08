package commands

import (
	"fmt"
	"log"

	"github.com/kelvins/passager/internal/models"
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

	getCmd.Flags().StringP("key", "k", "", "Key to encrypt/decrypt data using AES")

	return getCmd
}

func getCmdRun(cmd *cobra.Command, args []string) {
	key, err := cmd.Flags().GetString("key")
	if len(key) > 0 {
		if len(key) != 16 || err != nil {
			log.Fatal("Invalid key")
		} else {
			credential, err := models.Read(args[0])
			if err != nil {
				log.Fatal(err)
			}
			credential.Password = crypto.Decrypt(key, credential.Password)
			fmt.Printf("| %-24s| %-24s| %-24s|\n", "Name", "Login", "Password")
			fmt.Println(credential.String())
		}
	}
}

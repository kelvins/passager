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

	getCmd.Flags().StringP("key", "k", "", "Key to encrypt/decrypt data")

	return getCmd
}

func getCmdRun(cmd *cobra.Command, args []string) {
	credential, err := models.Read(args[0])
	if err != nil {
		log.Fatal(err)
	}
	key, err := cmd.Flags().GetString("key")
	if len(key) > 0 {
		if err != nil {
			log.Fatal("Invalid key")
		} else {
			credential.Password = crypto.Decrypt(credential.Password, key)
		}
	}
	fmt.Printf("| %-24s| %-24s| %-24s|\n", "Name", "Login", "Password")
	fmt.Println(credential.String())
}

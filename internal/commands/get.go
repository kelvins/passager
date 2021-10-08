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
	return getCmd
}

func getCmdRun(cmd *cobra.Command, args []string) {
	credential, err := models.Read(args[0])
	if err != nil {
		log.Fatal(err)
	}
	credential.Password = crypto.Decrypt("a1a2a3a4a5a6a7a8", credential.Password)
	fmt.Printf("| %-24s| %-24s| %-24s|\n", "Name", "Login", "Password")
	fmt.Println(credential.String())
}

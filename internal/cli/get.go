package cli

import (
	"fmt"

	"github.com/kelvins/passager/internal/db"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Set credentials to the database",
	Long:  "Set credentials to the database",
	Args:  cobra.MinimumNArgs(1),
	Run:   getCredential,
}

func getCredential(cmd *cobra.Command, args []string) {
	credential, err := db.Read(args[0])
	if err == nil {
		fmt.Printf("Login: %s\nPassword: %s\n", credential.Login, credential.Password)
	} else {
		fmt.Println("Not found")
	}
}

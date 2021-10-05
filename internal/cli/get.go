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
	Run:   getCredential,
}

func getCredential(cmd *cobra.Command, args []string) {
	credential, err := db.Read("test")
	if err == nil {
		fmt.Println(credential.Login)
	} else {
		fmt.Println("Not found")
	}
}

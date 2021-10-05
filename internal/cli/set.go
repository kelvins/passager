package cli

import (
	"github.com/kelvins/passager/internal/db"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set credentials to the database",
	Long:  "Set credentials to the database",
	Args:  cobra.MinimumNArgs(3),
	Run:   setCredential,
}

func setCredential(cmd *cobra.Command, args []string) {
	credential := db.Credential{
		Name:     args[0],
		Login:    args[1],
		Password: args[2],
	}
	db.Create(&credential)
}

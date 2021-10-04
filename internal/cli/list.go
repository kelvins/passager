package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Set credentials to the database",
	Long:  "Long description",
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Printf("List command")
	},
}

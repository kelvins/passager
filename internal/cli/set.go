package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set credentials to the database",
	Long:  "Long description",
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Printf("Set command")
	},
}

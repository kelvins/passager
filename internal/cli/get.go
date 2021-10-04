package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Set credentials to the database",
	Long:  "Long description",
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Printf("Get command")
	},
}

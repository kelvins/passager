package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "passenger",
	Short: "Your personal password manager",
	Long:  "A simple password manager",
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Printf("Root command")
	},
}

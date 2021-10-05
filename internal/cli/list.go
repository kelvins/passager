package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all database credentials",
	Long:  "List all database credentials",
	Run:   listCredentials,
}

func listCredentials(cmd *cobra.Command, args []string) {
	fmt.Println("run")
}

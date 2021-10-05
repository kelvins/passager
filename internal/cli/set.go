package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set credentials to the database",
	Long:  "Set credentials to the database",
	Run:   setCredential,
}

func setCredential(cmd *cobra.Command, args []string) {
	fmt.Println("run")
}

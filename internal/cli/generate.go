package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Randomly generate a secure password",
	Long:  "Long description",
	Run: generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	fmt.Println("new-password-123")
}

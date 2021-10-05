package cli

import (
	"fmt"
	"github.com/kelvins/passager/pkg/password"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Randomly generate a secure password",
	Long:  "Randomly generate a secure password",
	Run:   generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	fmt.Println(password.Generate(0))
	fmt.Println(cmd.Flag("length").Value)
}

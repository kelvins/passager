package cli

import (
	"fmt"
	"log"
	"github.com/kelvins/passager/pkg/password"
	"github.com/spf13/cobra"
)

func GenerateCmdFactory() *cobra.Command {
	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Randomly generate a secure password",
		Long:  "Randomly generate a secure password",
		Run:   generateCmdRun,
	}

	generateCmd.PersistentFlags().Int8P("length", "l", 12, "Password length. Min. 6 Max. 24")

	return generateCmd
}

func generateCmdRun(cmd *cobra.Command, args []string) {
	length, err := cmd.Flags().GetInt8("length")
	if length < 6 || length > 24 || err != nil {
		log.Fatal("Invalid length")
	}
	fmt.Println(password.Generate(0))
	fmt.Println(cmd.Flag("length").Value)
}

package commands

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

	generateCmd.PersistentFlags().Int8P("length", "l", 18, "Password length. Min. 6 Max. 32")
	generateCmd.Flags().BoolP("no-numbers", "", false, "Disable numbers on password")
	generateCmd.Flags().BoolP("no-symbols", "", false, "Disable symbols on password")

	return generateCmd
}

func generateCmdRun(cmd *cobra.Command, args []string) {
	length, err := cmd.Flags().GetInt8("length")
	if length < 6 || length > 32 || err != nil {
		log.Fatal("Invalid length")
	}
	noNumbers, _ := cmd.Flags().GetBool("no-numbers")
	noSymbols, _ := cmd.Flags().GetBool("no-symbols")
	fmt.Println(password.Generate(length, !noNumbers, !noSymbols))
}

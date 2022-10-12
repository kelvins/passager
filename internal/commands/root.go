package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "0.4.0"

// RootCmdFactory is responsible for creating the root passager cobra.Command.
func RootCmdFactory() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "passager",
		Short: "Your personal password manager.",
		Run:   rootCmdRun,
	}
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolP("version", "v", false, "Show CLI version")
	return rootCmd
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetBool("version")
	if version {
		fmt.Fprintf(cmd.OutOrStdout(), "Passager v%s\n", Version)
	}
}

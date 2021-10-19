package commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Version = "0.1.0"

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
		color.Green("Passager v%s", Version)
	}
}

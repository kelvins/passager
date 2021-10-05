package cli

import (
	"github.com/spf13/cobra"
)

func RootCmdFactory() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "passager",
		Short: "Your personal password manager",
		Long:  "A simple password manager",
	}

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
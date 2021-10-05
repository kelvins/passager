package main

import (
	"github.com/kelvins/passager/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "passager",
		Short: "Your personal password manager",
		Long:  "A simple password manager",
	}

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	cli.GenerateCmd.PersistentFlags().IntP("length", "l", 12, "Password length")

	rootCmd.AddCommand(cli.GenerateCmd)
	rootCmd.AddCommand(cli.SetCmd)
	rootCmd.AddCommand(cli.GetCmd)
	rootCmd.AddCommand(cli.ListCmd)

	rootCmd.Execute()
}

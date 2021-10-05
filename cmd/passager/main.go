package main

import (
	"github.com/kelvins/passager/internal/commands"
)

func main() {
	rootCmd := commands.RootCmdFactory()
	rootCmd.AddCommand(commands.GenerateCmdFactory())
	rootCmd.AddCommand(commands.SetCmdFactory())
	rootCmd.AddCommand(commands.GetCmdFactory())
	rootCmd.AddCommand(commands.ListCmdFactory())
	rootCmd.AddCommand(commands.DeleteCmdFactory())
	rootCmd.Execute()
}

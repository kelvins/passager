package main

import (
	"github.com/kelvins/passager/internal/cli"
)

func main() {
	rootCmd := cli.RootCmdFactory()
	rootCmd.AddCommand(cli.GenerateCmdFactory())
	rootCmd.AddCommand(cli.SetCmdFactory())
	rootCmd.AddCommand(cli.GetCmdFactory())
	rootCmd.AddCommand(cli.ListCmdFactory())
	rootCmd.Execute()
}

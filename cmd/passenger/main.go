package main

import "github.com/kelvins/passenger/internal/cli"

func main() {
	cli.RootCmd.AddCommand(cli.GenerateCmd)
	cli.RootCmd.AddCommand(cli.SetCmd)
	cli.RootCmd.AddCommand(cli.GetCmd)
	cli.RootCmd.AddCommand(cli.ListCmd)
	cli.RootCmd.Execute()
}

package main

import (
	"khidr/todo/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cmd.API()
	cobra.CheckErr(rootCmd.Execute())
}

package cmd

import (
	"os"

	"github.com/mohammadhsn/gonya/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "gonya",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.Load("config.yaml"); err != nil {
			panic("failed to load config.yaml: " + err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(httpCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

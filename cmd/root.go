package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version string = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vrcpic",
	Short:   "CLI tool for screenshots of VRChat",
	Long:    `vrcpic is a CLI tool for screenshots of VRChat.`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

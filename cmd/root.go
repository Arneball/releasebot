package cmd

import (
	"git.vgregion.se/healthapps/releasebot/botstuff"
	"git.vgregion.se/healthapps/releasebot/slowlane"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "releasebot",
	Short: "Assorted collection of useful commands",
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
	rootCmd.AddCommand(slowlane.UploadCommand, slowlane.NextVersionCommand, botstuff.TeamsCommand)
}

package cmd

import (
	"os"

	"github.com/Arneball/releasebot/botstuff"
	"github.com/Arneball/releasebot/slowlane"
	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(
		slowlane.UploadCommand,
		slowlane.NextVersionCommand,
		slowlane.UpdateTrack,
		botstuff.TeamsCommand,
		slowlane.DownloadApk,
	)
}

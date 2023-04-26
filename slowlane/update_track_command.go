package slowlane

import "github.com/spf13/cobra"

var UpdateTrack = &cobra.Command{
	Use:   "updateTrack",
	Short: "Update a track to point to a specific version",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()
		pkg, err := flags.GetString("package")
		if err != nil {
			return err
		}
		version, err := flags.GetInt64("version")
		if err != nil {
			return err
		}
		track, err := flags.GetString("track")
		if err != nil {
			return err
		}
		return AddToTrack(pkg, version, track)
	},
}

func init() {
	UpdateTrack.Flags().String("package", "", "android package name")
	UpdateTrack.Flags().Int64("version", 0, "aab version")
	UpdateTrack.Flags().String("track", "", "target track")
	for _, f := range []string{"package", "version", "track"} {
		if err := UpdateTrack.MarkFlagRequired(f); err != nil {
			panic(err)
		}
	}
}

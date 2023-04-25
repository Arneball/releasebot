package slowlane

import "github.com/spf13/cobra"

var UploadCommand = &cobra.Command{
	Use:   "upload",
	Short: "upload aab",
	Long:  `Uploads aab to google play, given --track, --package, --aab and --status`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		getOrPanic := func(key string) string {
			value, err := flags.GetString(key)
			if err != nil {
				panic(err)
			}
			return value
		}
		if err := Upload(getOrPanic("package"), getOrPanic("track"), getOrPanic("aab"), getOrPanic("status")); err != nil {
			panic(err)
		}
	},
}

func init() {
	UploadCommand.Flags().String("package", "", "android package name")
	UploadCommand.Flags().String("track", "", "track in google play")
	UploadCommand.Flags().String("aab", "", "aab path")
	UploadCommand.Flags().String("status", "", "whether the publication is draft or completed")
	for _, s := range []string{
		"package", "track", "aab", "status",
	} {
		err := UploadCommand.MarkFlagRequired(s)
		if err != nil {
			panic(err)
		}
	}
}

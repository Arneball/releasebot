package slowlane

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

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
	if err := UploadCommand.RegisterFlagCompletionFunc("package", ourPackages); err != nil {
		panic(err)
	}
	UploadCommand.Flags().String("track", "", "track in google play, comes from RELEASEBOT_APPS env.")
	UploadCommand.Flags().String("aab", "", "aab path")
	if err := UploadCommand.RegisterFlagCompletionFunc("aab", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return listFilesWithExtension(".", ".aab"), cobra.ShellCompDirectiveNoSpace
	}); err != nil {
		panic(err)
	}
	UploadCommand.Flags().String("status", "", "whether the publication is draft or completed")
	err := UploadCommand.RegisterFlagCompletionFunc("status", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"draft", "completed"}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		panic("lol?")
	}
	for _, s := range []string{
		"package", "track", "aab", "status",
	} {
		err := UploadCommand.MarkFlagRequired(s)
		if err != nil {
			panic(err)
		}
	}
}

// Helper function to list files recursively with a specific extension
func listFilesWithExtension(baseDir, extension string) []string {
	var files []string
	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == extension {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
		return nil
	}
	return files
}

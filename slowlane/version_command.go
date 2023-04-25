package slowlane

import (
	"fmt"
	"github.com/spf13/cobra"
)

var NextVersionCommand = &cobra.Command{
	Use:   "nextVersionCode",
	Short: "gets the next version code",
	Run: func(cmd *cobra.Command, args []string) {
		pkg, err := cmd.Flags().GetString("package")
		if err != nil {
			panic(err)
		}
		if code, err := GetNextVersion(pkg); err == nil {
			fmt.Printf("%d", code)
		} else {
			panic(err)
		}
	},
}

func init() {
	NextVersionCommand.Flags().String("package", "", "android package name")
	if err := NextVersionCommand.MarkFlagRequired("package"); err != nil {
		panic(err)
	}
}

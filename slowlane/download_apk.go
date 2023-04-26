package slowlane

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/api/androidpublisher/v3"
)

var DownloadApk = &cobra.Command{
	Use:   "downloadApk",
	Short: "Download an apk",
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
		return downloadApk(pkg, version)
	},
}

func init() {
	DownloadApk.Flags().String("package", "", "android package name")
	err := DownloadApk.RegisterFlagCompletionFunc("package", ourPackages)
	if err != nil {
		panic(err)
	}

	DownloadApk.Flags().Int64("version", 0, "aab version")
	for _, f := range []string{"package", "version"} {
		if err := DownloadApk.MarkFlagRequired(f); err != nil {
			panic(err)
		}
	}
}

func ourPackages(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	pkgs := []string{"com.vgregion.hud", "com.vgregion.migraine", "com.vgregion.epilepsy"}
	return pkgs, cobra.ShellCompDirectiveNoSpace
}

func downloadApk(pkg string, version int64) error {
	service, err := getService()
	if err != nil {
		return err
	}
	s := androidpublisher.NewGeneratedapksService(service)
	listResp, err := s.List(pkg, version).Do()
	if err != nil {
		return err
	}
	dlId := listResp.GeneratedApks[0].GeneratedUniversalApk.DownloadId
	dlCall := s.Download(pkg, version, dlId)
	resp, err := dlCall.Download()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	path := fmt.Sprintf("/tmp/%s_%d.apk", pkg, version)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	log.Printf("Downloaded %s to %s", pkg, path)
	return nil
}

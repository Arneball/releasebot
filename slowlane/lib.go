package slowlane

import (
	"context"
	"errors"
	android "google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"log"
	"os"
)

func Upload(pkg, track, file, state string) error {
	_, err := inEdit[bool](pkg, func(s *android.Service, edit *android.AppEdit) (bool, error) {
		f, err := os.Open(file)
		if err != nil {
			return false, err
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Printf("Closing failed: %s", err)
			}
		}()
		bundle, err := s.Edits.Bundles.Upload(pkg, edit.Id).
			Media(f, googleapi.ContentType("application/octet-stream")).
			Do()
		if err != nil {
			return false, err
		}
		_, err = s.Edits.Tracks.Update(pkg, edit.Id, track, &android.Track{
			Releases: []*android.TrackRelease{
				{
					VersionCodes: []int64{bundle.VersionCode},
					Status:       state,
				},
			},
		}).Do()
		if err != nil {
			return false, err
		}
		return true, nil
	})
	return err
}

func getService() (*android.Service, error) {
	getenv := os.Getenv("FASTLANE_KEY")
	if getenv == "" {
		return nil, errors.New("FASTLANE_KEY env empty")
	}
	return android.NewService(context.Background(), option.WithCredentialsFile(getenv))
}

func inEdit[T any](pkg string, f func(s *android.Service, edit *android.AppEdit) (T, error)) (T, error) {
	var def T
	s, err := getService()
	if err != nil {
		return def, err
	}
	edit, err := s.Edits.Insert(pkg, nil).Do()
	if err != nil {
		return def, err
	}
	result, err := f(s, edit)
	if err != nil {
		s.Edits.Delete(pkg, edit.Id)
		return result, err
	}
	_, err = s.Edits.Commit(pkg, edit.Id).Do()
	return result, err
}

var notReallyAnError = errors.New("notReallyAnError")

func GetNextVersion(pkg string) (int64, error) {
	cnt, err := inEdit[int64](pkg, func(s *android.Service, edit *android.AppEdit) (int64, error) {
		meh, err := s.Edits.Bundles.List(pkg, edit.Id).Do()
		if err != nil {
			return 0, err
		}
		var max = int64(0)
		for _, apk := range meh.Bundles {
			if apk.VersionCode > max {
				max = apk.VersionCode
			}
		}
		return max + 1, notReallyAnError
	})
	if err == notReallyAnError {
		err = nil
	}
	return cnt, err
}

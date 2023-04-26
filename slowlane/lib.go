package slowlane

import (
	"context"
	"errors"
	"log"
	"os"

	android "google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
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

func GetNextVersion(pkg string) (int64, error) {
	// Used so that inEdit discards the edit
	var discardEditError = errors.New("discardEditError")

	cnt, err := inEdit[int64](pkg, func(s *android.Service, edit *android.AppEdit) (int64, error) {
		meh, err := s.Edits.Bundles.List(pkg, edit.Id).Do()
		if err != nil {
			return 0, err
		}
		maxVersionCode := int64(0)
		for _, apk := range meh.Bundles {
			maxVersionCode = max(apk.VersionCode, maxVersionCode)
		}
		return maxVersionCode + 1, discardEditError
	})
	if errors.Is(err, discardEditError) {
		err = nil
	}
	return cnt, err
}

func AddToTrack(pkg string, version int64, track string) error {
	_, err := inEdit[struct{}](pkg, func(s *android.Service, edit *android.AppEdit) (struct{}, error) {
		_, err := s.Edits.Tracks.Update(pkg, edit.Id, track, &android.Track{
			Releases: []*android.TrackRelease{
				{
					VersionCodes: []int64{version},
					Status:       "completed",
				},
			},
		}).Do()
		return struct{}{}, err
	})
	return err
}

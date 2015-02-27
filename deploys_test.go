package empire

import (
	"testing"
)

func TestDeploysServiceDeploy(t *testing.T) {
	var released bool

	a := &mockAppsService{}
	c := &mockConfigsService{}
	s := &mockSlugsService{}
	r := &mockReleasesService{
		CreateFunc: func(app *App, config *Config, slug *Slug) (*Release, error) {
			released = true
			return nil, nil
		},
	}

	d := &deploysService{
		AppsService:     a,
		ConfigsService:  c,
		SlugsService:    s,
		ReleasesService: r,
	}

	image := Image{
		Repo: "remind101/r101-api",
		ID:   "1234",
	}

	if _, err := d.Deploy(image); err != nil {
		t.Fatal(err)
	}

	if got, want := released, true; got != want {
		t.Fatal("Expected a release to be created")
	}
}
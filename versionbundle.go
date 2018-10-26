package main

import (
	"errors"
)

type versionBundle struct {
	versions map[int]string
}

// NewMapVersionBundle creates a new in-memory version collection
func newMapVersionBundle() *versionBundle {
	bundle := &versionBundle{}
	bundle.versions = map[int]string
	return bundle
}

func (bundle *versionBundle) addVersion(version string) (err error) {
	versionCode = 0
	bundle.versions[versionCode] = version
	return err
}

func (bundle *versionBundle) getAllVersions() (versions []string, err error) {
	versions = make([]string, 0, len(bundle.versions))
	for _, value := range bundle.versions {
		versions = append(versions, value)
	}
	return versions, err
}

func (bundle *versionBundle) getVersion(code int) (version string, err error) {
	version, found := bundle[code]
	if !found {
		err = errors.New("Could not find version in bundle")
	}
	return version, err
}


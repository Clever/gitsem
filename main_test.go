package main

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/blang/semver.v3"
	"testing"
)

var bumpTests = []struct {
	Old          string
	New          string
	Change       string
	isPreRelease bool
}{
	{Old: "1.0.0", New: "2.0.0", Change: "major", isPreRelease: false},
	{Old: "1.1.0", New: "2.0.0", Change: "major", isPreRelease: false},
	{Old: "1.1.1", New: "2.0.0", Change: "major", isPreRelease: false},
	{Old: "1.0.0", New: "1.1.0", Change: "minor", isPreRelease: false},
	{Old: "1.0.1", New: "1.1.0", Change: "minor", isPreRelease: false},
	{Old: "1.1.1", New: "1.1.2", Change: "patch", isPreRelease: false},

	{Old: "1.0.0", New: "2.0.0-1", Change: "major", isPreRelease: true},
	{Old: "1.1.0", New: "2.0.0-1", Change: "major", isPreRelease: true},
	{Old: "1.1.1", New: "2.0.0-1", Change: "major", isPreRelease: true},
	{Old: "1.0.0", New: "1.1.0-1", Change: "minor", isPreRelease: true},
	{Old: "1.0.1", New: "1.1.0-1", Change: "minor", isPreRelease: true},
	{Old: "1.1.1", New: "1.1.2-1", Change: "patch", isPreRelease: true},
	{Old: "1.0.0-1", New: "1.0.0-2", Change: "major", isPreRelease: true},
	{Old: "1.1.1-1", New: "1.1.1-2", Change: "minor", isPreRelease: true},
	{Old: "1.1.1-1", New: "1.1.1-2", Change: "patch", isPreRelease: true},
}

func TestBump(t *testing.T) {
	for _, test := range bumpTests {
		v, err := semver.New(test.Old)
		assert.Nil(t, err)
		assert.Equal(t, test.New, bump(v, test.Change, test.isPreRelease).String())
	}
}

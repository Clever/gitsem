package main

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/blang/semver.v1"
	"testing"
)

var bumpTests = []struct {
	Old    string
	New    string
	Change string
}{
	{Old: "1.0.0", New: "2.0.0", Change: "major"},
	{Old: "1.1.0", New: "2.0.0", Change: "major"},
	{Old: "1.1.1", New: "2.0.0", Change: "major"},
	{Old: "1.0.0", New: "1.1.0", Change: "minor"},
	{Old: "1.0.1", New: "1.1.0", Change: "minor"},
	{Old: "1.1.1", New: "1.1.2", Change: "patch"},
}

func TestBump(t *testing.T) {
	for _, test := range bumpTests {
		v, err := semver.New(test.Old)
		assert.Nil(t, err)
		assert.Equal(t, test.New, bump(v, test.Change).String())
	}
}

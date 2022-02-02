package cmd_test

import (
	"github.com/Masterminds/semver/v3"
	"github.com/garethjevans/apachedist-resource/cmd"
	"github.com/garethjevans/apachedist-resource/download/downloadfakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckCmd_Run_InitialVersion(t *testing.T) {
	fake := &downloadfakes.FakeDownloader{}
	check := cmd.CheckCmd{
		Downloader: fake,
	}

	fake.GetVersionsReturns([]*semver.Version{
		semver.MustParse("8.0.1"),
		semver.MustParse("8.0.2"),
		semver.MustParse("8.0.3")}, nil)

	v, err := check.RunWithInput(cmd.In{
		Source: cmd.Source{
			Repository: "https://archive.apache.org/dist/tomcat/tomcat-8",
		},
	})

	t.Logf("versions = %s", v)
	assert.NoError(t, err)
	assert.Equal(t, len(v), 3)
	assert.Equal(t, fake.GetVersionsCallCount(), 1)
}

func TestCheckCmd_Run_AfterVersion(t *testing.T) {
	fake := &downloadfakes.FakeDownloader{}
	check := cmd.CheckCmd{
		Downloader: fake,
	}

	fake.GetVersionsReturns([]*semver.Version{
		semver.MustParse("8.0.1"),
		semver.MustParse("8.0.2"),
		semver.MustParse("8.0.3")}, nil)

	v, err := check.RunWithInput(cmd.In{
		Source: cmd.Source{
			Repository: "https://archive.apache.org/dist/tomcat/tomcat-8",
		},
		Version: cmd.Version{Ref: "8.0.2"},
	})

	t.Logf("versions = %s", v)
	assert.NoError(t, err)
	assert.Equal(t, len(v), 2)
	assert.Equal(t, fake.GetVersionsCallCount(), 1)
}

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/garethjevans/apachedist-resource/download"
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var semverRE = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

type CheckCmd struct {
	Command    *cobra.Command
	Downloader download.Downloader
}

func NewCheckCmd() CheckCmd {
	check := CheckCmd{
		Downloader: &download.DefaultDownloader{},
	}
	check.Command = &cobra.Command{
		Use:   "check",
		Short: "checks a resource",
		Long:  `checks a resource`,
		Run:   check.Run,
	}

	return check
}

func (i *CheckCmd) RunWithInput(jsonIn In) ([]Version, error) {
	Log("Checking resource for %+v\n", jsonIn)

	var err error
	var versionToCheck *semver.Version
	if jsonIn.Version.Ref != "" {
		versionToCheck, err = semver.NewVersion(jsonIn.Version.Ref)
		if err != nil {
			Log("Skipping existing version %+s, %s\n", jsonIn.Version.Ref, err)
		}
	}

	versions, err := i.Downloader.GetVersions(jsonIn.Source.Repository)
	if err != nil {
		return nil, err
	}

	Log("got %d versions, filtering...\n", len(versions))
	var refs []Version
	for _, version := range versions {
		if versionToCheck == nil || *versionToCheck == *version || versionToCheck.LessThan(version) {
			refs = append(refs, Version{Ref: version.String()})
		}
	}
	Log("returning %s\n", refs)
	return refs, nil
}

func (i *CheckCmd) Run(cmd *cobra.Command, args []string) {
	var jsonIn In

	err := json.NewDecoder(os.Stdin).Decode(&jsonIn)
	if err != nil {
		log.Fatal(err)
	}

	v, err := i.RunWithInput(jsonIn)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

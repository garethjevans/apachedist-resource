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
	Command *cobra.Command
}

func NewCheckCmd() CheckCmd {
	check := CheckCmd{}
	check.Command = &cobra.Command{
		Use:   "check",
		Short: "checks a resource",
		Long:  `checks a resource`,
		Run:   check.Run,
	}

	return check
}

func (i *CheckCmd) Run(cmd *cobra.Command, args []string) {
	var jsonIn In

	err := json.NewDecoder(os.Stdin).Decode(&jsonIn)
	if err != nil {
		log.Fatal(err)
	}

	Log("Checking resource for %+v\n", jsonIn)

	var versionToCheck *semver.Version
	if jsonIn.Version.Ref != "" {
		versionToCheck, err = semver.NewVersion(jsonIn.Version.Ref)
		if err != nil {
			Log("Skipping existing version %+s, %s\n", jsonIn.Version.Ref, err)
		}
	}

	versions, err := download.GetVersions(jsonIn.Source.Repository)
	if err != nil {
		panic(err)
	}

	Log("got %d versions, filtering...\n", len(versions))
	var refs []Version
	for _, version := range versions {
		if versionToCheck == nil || *versionToCheck == *version || versionToCheck.LessThan(version) {
			refs = append(refs, Version{Ref: version.String()})
		}
	}
	Log("returning %s\n", refs)
	b, err := json.Marshal(refs)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

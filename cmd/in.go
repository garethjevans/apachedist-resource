package cmd

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/garethjevans/apachedist-resource/download"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type InCmd struct {
	Command    *cobra.Command
	Downloader download.Downloader
}

func NewInCmd() InCmd {
	in := InCmd{
		Downloader: &download.DefaultDownloader{},
	}
	in.Command = &cobra.Command{
		Use:   "in",
		Short: "concourse in command",
		Long:  `concourse in command`,
		Run:   in.Run,
	}

	return in
}

func (i *InCmd) Run(cmd *cobra.Command, args []string) {
	var jsonIn In

	err := json.NewDecoder(os.Stdin).Decode(&jsonIn)
	if err != nil {
		log.Fatal(err)
	}

	outputDir := args[0]

	artifact, err := i.Downloader.Download(jsonIn.Source.ArtifactId, jsonIn.Version.Ref,
		outputDir, jsonIn.Source.Repository, jsonIn.Source.Extension)

	if err != nil {
		panic(err)
	}

	// lets validate sha1, this should always exist
	downloadedFilePath := path.Join(outputDir, artifact.Filename)
	downloadedFileContents, err := ioutil.ReadFile(downloadedFilePath)
	if err != nil {
		panic(err)
	}

	sha512 := sha512.Sum512(downloadedFileContents)

	if fmt.Sprintf("%x", sha512) != artifact.Sha512 {
		log.Fatalf("calculated sha512 does not match downloaded sha512: %x != %s\n", sha512, artifact.Sha512)
	} else {
		Log("sha512 %s is valid\n", artifact.Sha512)
	}

	// if Sha256 does exist, calculate it
	if artifact.Sha256 == "" {
		Log("sha256 does not exist, calculating it from downloaded file\n")
		sha256 := sha256.Sum256(downloadedFileContents)
		artifact.Sha256 = fmt.Sprintf("%x", sha256)
	}

	out := InResponse{
		Version: Version{Ref: artifact.Version},
		Metadata: []Metadata{
			{Name: "version", Value: artifact.Version},
			{Name: "uri", Value: artifact.Url},
			{Name: "filename", Value: artifact.Filename},
			{Name: "cpe", Value: fmt.Sprintf("cpe:2.3:a:%s:%s:%s:*:*:*:*:*:*:*", jsonIn.Source.ArtifactId, jsonIn.Source.ArtifactId, artifact.Version)},
			{Name: "sha512", Value: artifact.Sha512},
			{Name: "sha256", Value: artifact.Sha256},
		},
	}

	for _, m := range out.Metadata {
		file := path.Join(outputDir, m.Name)
		Log("creating %s\n", file)
		err = ioutil.WriteFile(file, []byte(m.Value), 0644)
		if err != nil {
			panic(err)
		}
	}

	b, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

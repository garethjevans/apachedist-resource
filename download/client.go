package download

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Downloader

type Artifact struct {
	Version       string
	RepositoryUrl string
	UrlPatten     string
	Downloader    func(string, string, string) (*http.Response, error)
}

type DownloadedArtifact struct {
	Filename string
	Url      string
	Version  string
	Sha256   string
	Sha512   string
}

var semverRE = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
var semverFinder = regexp.MustCompile(`(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)

type Downloader interface {
	GetVersions(repository string) ([]*semver.Version, error)
	Download(version, dest, repo, urlPattern string) (*DownloadedArtifact, error)
}

type DefaultDownloader struct {
}

func (d *DefaultDownloader) GetVersions(repository string) ([]*semver.Version, error) {
	a := Artifact{
		RepositoryUrl: repository,
		Downloader:    httpGetCustom,
	}

	v, err := AllVersions(a)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (d *DefaultDownloader) Download(version, dest, repo, urlPattern string) (*DownloadedArtifact, error) {
	a := Artifact{
		UrlPatten:     urlPattern,
		Version:       version,
		RepositoryUrl: repo,
		Downloader:    httpGetCustom,
	}

	return DownloadArtifact(a, dest)
}

func DownloadArtifact(a Artifact, dest string) (*DownloadedArtifact, error) {
	url := ArtifactUrl(a)
	//fmt.Println("DEBUG", url)
	resp, err := a.Downloader(url, "", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	filename := FileName(a)
	filepath := dest + "/" + filename

	out, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}

	sha512 := Sha512(a)
	sha256 := Sha256(a)

	return &DownloadedArtifact{Version: a.Version, Url: url, Filename: filename, Sha512: sha512, Sha256: sha256}, nil
}

func httpGetCustom(url, user, pwd string) (*http.Response, error) {
	if user != "" && pwd != "" {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth(user, pwd)
		return client.Do(req)
	}

	return http.Get(url)
}

func FileName(a Artifact) string {
	url := ArtifactUrl(a)
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func ArtifactUrl(a Artifact) string {
	// https://archive.apache.org/dist/tomee/tomee-9.0.0-M7/apache-tomee-9.0.0-M7-microprofile.tar.gz
	// https://archive.apache.org/dist/tomee/tomcat-9/9.0.0-M7/bin/apache-tomcat-9.0.0-M7.tar.gz
	// FIXME should ensure that repo url has a trailing slash
	path := strings.ReplaceAll(a.UrlPatten, "${version}", a.Version)
	return a.RepositoryUrl + path
}

func Sha512(a Artifact) string {
	url := ArtifactUrl(a) + ".sha512"
	r, err := a.Downloader(url, "", "")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return ""
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(body), " ")[0]
}

func Sha256(a Artifact) string {
	url := ArtifactUrl(a) + ".sha256"
	r, err := a.Downloader(url, "", "")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return ""
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(body), " ")[0]
}

func AllVersions(a Artifact) ([]*semver.Version, error) {
	// FIXME should ensure that repo url has a trailing slash
	metadataUrl := a.RepositoryUrl
	resp, err := a.Downloader(metadataUrl, "", "")
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unable to fetch maven metadata from %s Http statusCode: %d", metadataUrl, resp.StatusCode)
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var versions []string
	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if semverFinder.MatchString(s.Text()) {
			version := semverFinder.FindStringSubmatch(s.Text())[0]
			// fmt.Println("DEBUG", semverFinder.FindStringSubmatch(s.Text()))
			// first we check if the version doesn't contain any extra characters removed from the string
			if strings.HasSuffix(s.Text(), fmt.Sprintf("%s/", version)) {
				// lets use the strict check
				if semverRE.MatchString(version) {
					versions = append(versions, version)
				}
			}
		}
	})

	vs := make([]*semver.Version, len(versions))
	for i, r := range versions {
		v := semver.MustParse(r)
		vs[i] = v
	}

	sort.Sort(semver.Collection(vs))

	//fmt.Println(vs)

	return vs, nil
}

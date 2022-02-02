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
	ArtifactId    string
	Extension     string
	Version       string
	RepositoryUrl string
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

type Downloader interface {
	GetVersions(repository string) ([]*semver.Version, error)
	Download(artifactId, version, dest, repo, extension string) (*DownloadedArtifact, error)
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

func (d *DefaultDownloader) Download(artifactId, version, dest, repo, extension string) (*DownloadedArtifact, error) {
	a := Artifact{
		ArtifactId:    artifactId,
		Extension:     extension,
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
	return fmt.Sprintf("%s-%s.%s", a.ArtifactId, a.Version, a.Extension)
}

func ArtifactUrl(a Artifact) string {
	// FIXME should ensure that repo url has a trailing slash
	return a.RepositoryUrl + "/v" + a.Version + "/bin/" + FileName(a)
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
		if strings.HasPrefix(s.Text(), "v") {
			v := strings.TrimSuffix(strings.TrimPrefix(s.Text(), "v"), "/")
			if semverRE.MatchString(v) {
				versions = append(versions, v)
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

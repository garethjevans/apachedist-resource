package download

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestAllVersions_Tomee(t *testing.T) {
	a := Artifact{
		RepositoryUrl: "https://archive.apache.org/dist/tomee",
		Downloader: func(url string, user string, password string) (*http.Response, error) {
			assert.Equal(t, url, "https://archive.apache.org/dist/tomee")

			r := &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(strings.NewReader(`
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<html>
 <head>
  <title>Index of /dist/tomee</title>
 </head>
 <body>
<h1>Index of /dist/tomee</h1>
<pre><img src="/icons/blank.gif" alt="Icon "> <a href="?C=N;O=D">Name</a>                    <a href="?C=M;O=A">Last modified</a>      <a href="?C=S;O=A">Size</a>  <a href="?C=D;O=A">Description</a><hr><img src="/icons/back.gif" alt="[PARENTDIR]"> <a href="/dist/">Parent Directory</a>                             -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.5.2/">tomee-1.5.2/</a>            2013-04-06 10:51    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.6.0.1/">tomee-1.6.0.1/</a>          2014-04-19 19:06    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.6.0.2/">tomee-1.6.0.2/</a>          2020-07-03 04:07    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.6.0/">tomee-1.6.0/</a>            2013-11-20 10:09    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.7.0/">tomee-1.7.0/</a>            2015-02-17 20:04    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.7.1/">tomee-1.7.1/</a>            2015-02-17 20:04    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.7.2/">tomee-1.7.2/</a>            2015-05-22 15:45    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.7.3/">tomee-1.7.3/</a>            2015-12-09 11:30    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.7.4/">tomee-1.7.4/</a>            2017-10-04 10:58    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-1.7.5/">tomee-1.7.5/</a>            2020-07-03 04:08    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.0-M1/">tomee-7.0.0-M1/</a>         2015-12-11 21:02    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.0-M2/">tomee-7.0.0-M2/</a>         2016-03-09 08:33    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.0-M3/">tomee-7.0.0-M3/</a>         2016-03-06 18:12    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.0/">tomee-7.0.0/</a>            2016-05-29 15:49    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.1/">tomee-7.0.1/</a>            2016-06-27 11:01    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.2/">tomee-7.0.2/</a>            2016-11-11 19:04    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.3/">tomee-7.0.3/</a>            2017-10-04 10:58    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.4/">tomee-7.0.4/</a>            2018-05-04 19:50    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.5/">tomee-7.0.5/</a>            2018-07-24 14:21    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.6/">tomee-7.0.6/</a>            2020-07-03 04:06    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.7/">tomee-7.0.7/</a>            2020-07-03 04:04    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.8/">tomee-7.0.8/</a>            2020-07-03 04:03    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.0.9/">tomee-7.0.9/</a>            2020-11-05 18:49    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.1.0/">tomee-7.1.0/</a>            2018-09-07 08:49    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.1.1/">tomee-7.1.1/</a>            2020-07-03 04:05    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.1.2/">tomee-7.1.2/</a>            2020-07-03 04:04    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.1.3/">tomee-7.1.3/</a>            2020-07-03 04:03    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-7.1.4/">tomee-7.1.4/</a>            2020-11-05 18:50    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.0-M1/">tomee-8.0.0-M1/</a>         2020-07-03 04:07    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.0-M2/">tomee-8.0.0-M2/</a>         2020-07-03 04:06    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.0-M3/">tomee-8.0.0-M3/</a>         2020-07-03 04:06    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.0/">tomee-8.0.0/</a>            2020-07-03 04:05    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.1/">tomee-8.0.1/</a>            2020-07-03 04:05    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.2/">tomee-8.0.2/</a>            2020-07-03 04:04    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.3/">tomee-8.0.3/</a>            2020-07-03 04:03    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.4/">tomee-8.0.4/</a>            2020-11-05 18:50    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.5/">tomee-8.0.5/</a>            2020-11-25 00:14    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.6/">tomee-8.0.6/</a>            2021-01-25 16:17    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.7/">tomee-8.0.7/</a>            2021-05-15 03:24    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.8/">tomee-8.0.8/</a>            2021-09-13 08:17    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-8.0.9/">tomee-8.0.9/</a>            2022-01-10 11:43    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-9.0.0-M2/">tomee-9.0.0-M2/</a>         2020-11-05 18:50    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-9.0.0-M3/">tomee-9.0.0-M3/</a>         2020-11-25 00:15    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-9.0.0-M7/">tomee-9.0.0-M7/</a>         2021-05-15 03:25    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="tomee-patch-plugin-0.8/">tomee-patch-plugin-0.8/</a> 2022-01-21 08:31    -   
<img src="/icons/unknown.gif" alt="[   ]"> <a href="KEYS">KEYS</a>                    2021-09-15 14:28   59K  
<hr></pre>
</body></html>
`)),
			}

			return r, nil
		},
	}
	allVersions, err := AllVersions(a)
	assert.NoError(t, err)
	assert.Equal(t, len(allVersions), 42)
	assert.Equal(t, semver.MustParse("1.5.2"), allVersions[0])
	assert.Equal(t, semver.MustParse("9.0.0-M7"), allVersions[41])
}

func TestAllVersions_Tomcat(t *testing.T) {
	a := Artifact{
		RepositoryUrl: "https://archive.apache.org/dist/tomcat/tomcat-9",
		Downloader: func(url string, user string, password string) (*http.Response, error) {
			assert.Equal(t, url, "https://archive.apache.org/dist/tomcat/tomcat-9")

			r := &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(strings.NewReader(`
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<html>
 <head>
  <title>Index of /dist/tomcat/tomcat-9</title>
 </head>
 <body>
<h1>Index of /dist/tomcat/tomcat-9</h1>
<pre><img src="/icons/blank.gif" alt="Icon "> <a href="?C=N;O=D">Name</a>                    <a href="?C=M;O=A">Last modified</a>      <a href="?C=S;O=A">Size</a>  <a href="?C=D;O=A">Description</a><hr><img src="/icons/back.gif" alt="[PARENTDIR]"> <a href="/dist/tomcat/">Parent Directory</a>                             -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M1/">v9.0.0.M1/</a>              2015-11-19 12:30    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M10/">v9.0.0.M10/</a>             2016-09-05 08:28    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M11/">v9.0.0.M11/</a>             2016-10-10 13:34    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M13/">v9.0.0.M13/</a>             2016-11-08 10:09    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M15/">v9.0.0.M15/</a>             2016-12-08 20:51    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M17/">v9.0.0.M17/</a>             2017-01-16 20:15    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M18/">v9.0.0.M18/</a>             2017-03-13 17:49    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M19/">v9.0.0.M19/</a>             2017-03-30 16:00    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M20/">v9.0.0.M20/</a>             2017-04-18 10:12    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M21/">v9.0.0.M21/</a>             2017-06-20 09:44    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M22/">v9.0.0.M22/</a>             2017-06-26 18:02    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M25/">v9.0.0.M25/</a>             2017-07-28 12:07    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M26/">v9.0.0.M26/</a>             2017-08-08 19:36    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M27/">v9.0.0.M27/</a>             2017-09-19 20:32    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M3/">v9.0.0.M3/</a>              2016-02-05 14:42    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M4/">v9.0.0.M4/</a>              2016-03-16 11:44    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M6/">v9.0.0.M6/</a>              2016-05-16 10:39    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M8/">v9.0.0.M8/</a>              2016-06-13 15:31    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.0.M9/">v9.0.0.M9/</a>              2016-07-12 19:35    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.1/">v9.0.1/</a>                 2017-10-04 11:11    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.10/">v9.0.10/</a>                2018-06-26 08:47    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.11/">v9.0.11/</a>                2018-08-17 13:14    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.12/">v9.0.12/</a>                2018-09-10 08:48    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.13/">v9.0.13/</a>                2018-11-07 23:05    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.14/">v9.0.14/</a>                2018-12-12 08:40    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.16/">v9.0.16/</a>                2019-09-06 07:09    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.17/">v9.0.17/</a>                2019-03-18 20:48    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.19/">v9.0.19/</a>                2019-04-13 22:28    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.2/">v9.0.2/</a>                 2017-11-30 08:50    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.20/">v9.0.20/</a>                2019-05-13 08:00    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.21/">v9.0.21/</a>                2019-06-07 20:25    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.22/">v9.0.22/</a>                2019-07-09 14:09    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.24/">v9.0.24/</a>                2019-08-17 16:35    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.26/">v9.0.26/</a>                2019-09-19 19:55    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.27/">v9.0.27/</a>                2019-10-11 08:46    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.29/">v9.0.29/</a>                2019-11-21 09:11    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.30/">v9.0.30/</a>                2019-12-12 08:22    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.31/">v9.0.31/</a>                2020-03-10 23:16    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.33/">v9.0.33/</a>                2020-03-16 09:37    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.34/">v9.0.34/</a>                2020-04-08 15:43    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.35/">v9.0.35/</a>                2020-05-11 11:19    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.36/">v9.0.36/</a>                2020-06-07 19:01    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.37/">v9.0.37/</a>                2020-07-05 18:31    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.38/">v9.0.38/</a>                2020-11-05 18:47    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.39/">v9.0.39/</a>                2020-11-05 18:47    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.4/">v9.0.4/</a>                 2018-01-22 14:51    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.40/">v9.0.40/</a>                2020-11-17 14:23    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.41/">v9.0.41/</a>                2020-12-08 17:34    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.43/">v9.0.43/</a>                2021-02-02 19:18    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.44/">v9.0.44/</a>                2021-03-10 11:31    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.45/">v9.0.45/</a>                2021-04-06 11:11    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.46/">v9.0.46/</a>                2021-05-12 21:55    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.48/">v9.0.48/</a>                2021-06-15 11:28    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.5/">v9.0.5/</a>                 2018-02-11 20:11    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.50/">v9.0.50/</a>                2021-07-02 13:59    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.52/">v9.0.52/</a>                2021-08-06 05:51    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.53/">v9.0.53/</a>                2021-09-10 07:18    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.54/">v9.0.54/</a>                2021-10-01 20:42    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.55/">v9.0.55/</a>                2021-11-15 19:26    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.56/">v9.0.56/</a>                2021-12-08 09:03    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.58/">v9.0.58/</a>                2022-01-20 13:53    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.6/">v9.0.6/</a>                 2018-03-08 18:55    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.7/">v9.0.7/</a>                 2018-05-04 16:40    -   
<img src="/icons/folder.gif" alt="[DIR]"> <a href="v9.0.8/">v9.0.8/</a>                 2018-05-04 16:40    -   
<img src="/icons/unknown.gif" alt="[   ]"> <a href="KEYS">KEYS</a>                    2021-06-15 12:16   14K  
<hr></pre>
</body></html>
`)),
			}

			return r, nil
		},
	}
	allVersions, err := AllVersions(a)
	assert.NoError(t, err)
	assert.Equal(t, len(allVersions), 45)
	assert.Equal(t, semver.MustParse("9.0.1"), allVersions[0])
	assert.Equal(t, semver.MustParse("9.0.58"), allVersions[44])
}

func TestDownloadArtifact(t *testing.T) {
	a := Artifact{
		Version:       "0.1",
		UrlPatten:     "/v${version}/bin/bong-${version}.zip.gz",
		RepositoryUrl: "https://archive.apache.org/dist/tomcat/tomcat-9",
		Downloader: func(url string, user string, password string) (*http.Response, error) {
			switch url {
			case "https://archive.apache.org/dist/tomcat/tomcat-9/v0.1/bin/bong-0.1.zip.gz":
				r := &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(`some dummy content`)),
				}
				return r, nil
			case "https://archive.apache.org/dist/tomcat/tomcat-9/v0.1/bin/bong-0.1.zip.gz.sha256":
				r := &http.Response{
					StatusCode: 404,
					Body:       ioutil.NopCloser(strings.NewReader(`404`)),
				}
				return r, nil
			case "https://archive.apache.org/dist/tomcat/tomcat-9/v0.1/bin/bong-0.1.zip.gz.sha512":
				r := &http.Response{
					StatusCode: 404,
					Body:       ioutil.NopCloser(strings.NewReader(`sha512 hash`)),
				}
				return r, nil
			}

			return nil, fmt.Errorf("unknown url %s", url)
		},
	}

	dir, err := ioutil.TempDir(".", "test-output")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	_, err = DownloadArtifact(a, dir)
	assert.NoError(t, err)
}

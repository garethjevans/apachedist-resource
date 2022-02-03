# Apache Distribution Resource

[![Go Report Card](https://goreportcard.com/badge/github.com/garethjevans/apachedist-resource)](https://goreportcard.com/report/github.com/garethjevans/apachedist-resource)

A concourse resource that can track information about an apache dist release, e.g. something hosted at https://archive.apache.org/dist

## Source Configuration

* `repository`: *Required.* The location of the repository.

* `url_pattern`: *Required.* The pattern of the url used to access the artifact.
  `${version}` can be used as a placeholder instead of the determined version. 
  e.g. for tomcat use `/v${version}/bin/apache-tomcat-${version}.tar.gz`

## Behavior

### `check`: Check for new versions of the artifact.

Checks for new versions of the artifact by retrieving the directory listing from 
the main repository page and parsing the HTML, attempting to extract an urls that
contain valid semantic versions. This can be a little fragile.

NOTE: Only valid semantic versions are supported.

### `in`: Fetch an artifact from a repository.

Download the artifact from the repository.

#### Additional files populated

* `version`: The version of the downloaded artifact.

* `filename`: The filename of the downloaded file.

* `uri`: The full uri that can be used to reference the current version of the artifact.

* `sha256`: The sha256 sum of the downloaded file. If this is not available from the maven 
   repository, it is calculated from the downloaded file after the `sha512` is validated.

* `sha512`: The sha512 sum of the downloaded file. 

* `cpe`: The version to be used in a CPE referrence.

* `purl`: The version to be used in a PURL reference.
 
### `out`:

Not Implemented.

## Examples

Resource configuration for new apache tomcat versions:

``` yaml
resource_types:
- name: apachedist-resource
  type: registry-image
  source:
    repository: ghcr.io/garethjevans/apachedist-resource
    tag: latest

resources:
- name: artifact
  type: apachedist-resource
  source:
    repository: https://archive.apache.org/dist/tomcat/tomcat-9
    url_pattern: /v${version}/bin/apache-tomcat-${version}.tar.gz
```

Resource configuration for new apache tomee versions:

``` yaml
resource_types:
- name: apachedist-resource
  type: registry-image
  source:
    repository: ghcr.io/garethjevans/apachedist-resource
    tag: latest

resources:
- name: artifact
  type: apachedist-resource
  source:
    repository: https://archive.apache.org/dist/tomee/
    url_pattern: /tomee-${version}/apache-tomee-${version}-webprofile.tar.gz
```

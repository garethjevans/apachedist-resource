package cmd

type Source struct {
	ArtifactId string `json:"artifactId"`
	Extension  string `json:"extension"`
	Repository string `json:"repository"`
	UrlPattern string `json:"url_pattern"`
}

type Version struct {
	Ref string `json:"ref"`
}

type In struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

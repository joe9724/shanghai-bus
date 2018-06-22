package models

type Version struct {
	Number     string `json:"number"`
	Content    string `json:"content"`
	Download   string `json:"download"`
	Forced     bool   `json:"forced"`
	VersionUrl string `json:"version_url"`
	AboutUrl   string `json:"about_url"`
}

type VersionData struct {
	Version Version `json:"version"`
	Splash  Splash  `json:"splash"`
}

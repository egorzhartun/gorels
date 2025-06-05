package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	repoOwner      = "egorzhartun"
	repoName       = "gorels"
	currentVersion = "v1.0.0" // обновляй при каждом релизе
)

type githubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

func CheckForUpdate() (bool, string, string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	resp, err := http.Get(url)
	if err != nil {
		return false, "", "", err
	}
	defer resp.Body.Close()

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return false, "", "", err
	}

	if release.TagName != currentVersion {
		return true, release.TagName, release.HTMLURL, nil
	}
	return false, "", "", nil
}

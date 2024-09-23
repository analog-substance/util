package updater

import (
	"encoding/json"
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strings"
	"syscall"
	"time"
)

func SelfUpdate(check bool) {
	buildInfo, ok := debug.ReadBuildInfo()

	if ok {
		mod := buildInfo.Main.Path
		modURL, err := url.Parse(fmt.Sprintf("https://%s", mod))
		if err != nil {
			log.Fatal(err)
		}

		if modURL.Host == "github.com" {
			gitHubAPIURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", modURL.Path[1:])
			resp, err := http.Get(gitHubAPIURL)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			var releaseInfo GitHubReleases
			err = json.Unmarshal(body, &releaseInfo)
			if err != nil {
				log.Fatal(err)
			}

			if isNewerVersion(buildInfo.Main.Version, releaseInfo.Name) {
				if check {
					fmt.Printf("%s is outdated. Current release version: %s\n", buildInfo.Main.Version, releaseInfo.Name)
				} else {
					err = goInstall(fmt.Sprintf("%s@latest", buildInfo.Main.Path))
					if err != nil {
						log.Fatal(err)
					}
				}
			} else {
				fmt.Printf("%s  is up to date. Current release version: %s\n", buildInfo.Main.Version, releaseInfo.Name)
			}
		} else {
			log.Panic("Unable to autoupdate non-github based projects.")
		}
	}
}

func goInstall(goRepo string) error {
	sshPath, _ := exec.LookPath("go")

	args := []string{
		"go",
		"install",
		goRepo,
	}

	//args = append(args, additionalArgs...)
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		return builder.Cmd(args[0], args[1:]...).Interactive().Run()
	}
	return syscall.Exec(sshPath, args, os.Environ())
}

type GitHubReleases struct {
	Url       string `json:"url"`
	AssetsUrl string `json:"assets_url"`
	UploadUrl string `json:"upload_url"`
	HtmlUrl   string `json:"html_url"`
	Id        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeId          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []struct {
		Url      string `json:"url"`
		Id       int    `json:"id"`
		NodeId   string `json:"node_id"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadUrl string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballUrl string `json:"tarball_url"`
	ZipballUrl string `json:"zipball_url"`
	Body       string `json:"body"`
}

func isNewerVersion(currentVersion, versionToCompare string) bool {
	currentVer := strings.Split(strings.TrimPrefix(currentVersion, "v"), ".")
	compareVer := strings.Split(strings.TrimPrefix(versionToCompare, "v"), ".")

	for i, v := range currentVer {
		if compareVer[i] > v {
			return true
		} else if compareVer[i] < v {
			return false
		}
	}

	return false
}

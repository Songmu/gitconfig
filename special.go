package gitconfig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

func (c *Config) User() (string, error) {
	return c.Get("user.name")
}

func (c *Config) Email() (string, error) {
	return c.Get("user.email")
}

func (c *Config) GitHubToken() (string, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		return token, nil
	}
	return c.Get("github.token")
}

func (c *Config) GitHubUser(host string) (string, error) {
	if host == "" {
		host = os.Getenv("GITHUB_HOST")
		if host == "" {
			host = "github.com"
		}
	}
	if user := os.Getenv("GITHUB_USER"); user != "" {
		return user, nil
	}
	if user, err := c.Get(fmt.Sprintf("credential.https://%s.username", host)); err == nil {
		return user, nil
	}
	if user, err := c.Get("github.user"); err == nil {
		return user, nil
	}
	if user, err := getGHUserFromHub(host); err == nil {
		return user, nil
	}
	if email, err := c.Email(); err == nil {
		apiHost := os.Getenv("GITHUB_API")
		if apiHost == "" {
			apiHost = host
		}
		if apiHost == "github.com" {
			apiHost = "api.github.com"
		}
		if user, err := getGHUserFromGHAPI(apiHost, email); err == nil {
			return user, nil
		}
	}
	return c.Get("user.username")
}

func getGHUserFromHub(host string) (string, error) {
	// XXX parsing ${XDG_CONFIG_HOME:.config}/hub is better?
	if _, err := exec.LookPath("hub"); err != nil {
		return "", err
	}
	cmd := exec.Command("hub", "api", "user")
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("GITHUB_HOST=%s", host))
	var s struct {
		Login string
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}
	if err := json.NewDecoder(buf).Decode(&s); err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return s.Login, nil
}

func getGHUserFromGHAPI(apiHost, email string) (string, error) {
	v := url.Values{}
	v.Add("q", fmt.Sprintf("%s in:email", email))
	v.Add("per_page", "2")
	u := &url.URL{
		Scheme:   "https",
		Host:     apiHost,
		Path:     "/search/users",
		RawQuery: v.Encode(),
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var s struct {
		TotalCount int `json:"total_count"`
		Items      []struct {
			Login string
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return "", err
	}
	switch s.TotalCount {
	case 0:
		return "", fmt.Errorf("no users found from GitHub")
	case 1:
		return s.Items[0].Login, nil
	}
	return getGHUserFromGHCommit(apiHost, email)
}

func getGHUserFromGHCommit(apiHost, email string) (string, error) {
	v := url.Values{}
	v.Add("q", fmt.Sprintf("author-email:%s", email))
	v.Add("sort", "author-date")
	v.Add("per_page", "1")
	u := &url.URL{
		Scheme:   "https",
		Host:     apiHost,
		Path:     "/search/commits",
		RawQuery: v.Encode(),
	}
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Add("Accept", "application/vnd.github.cloak-preview")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var s struct {
		TotalCount int `json:"total_count"`
		Items      []struct {
			Author struct {
				Login string
			}
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return "", err
	}
	if s.TotalCount < 1 {
		return "", fmt.Errorf("no commits found")
	}
	return s.Items[0].Author.Login, nil
}

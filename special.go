package gitconfig

import "os"

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

package gitconfig

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Config is for base setting for git config
type Config struct {
	System, Global, Local bool
	File                  string
	Cd                    string
	GitPath               string
}

func (c *Config) git() string {
	if c.GitPath != "" {
		return c.GitPath
	}
	return "git"
}

// Do the git config
func (c *Config) Do(args ...string) (string, error) {
	gitArgs := append([]string{"config", "--null"})
	if c.Cd != "" {
		gitArgs = append([]string{"-C", c.Cd}, gitArgs...)
	}
	if c.System {
		gitArgs = append(gitArgs, "--system")
	}
	if c.Global {
		gitArgs = append(gitArgs, "--global")
	}
	if c.Local {
		gitArgs = append(gitArgs, "--local")
	}
	if c.File != "" {
		gitArgs = append(gitArgs, "--file", c.File)
	}

	gitArgs = append(gitArgs, args...)
	cmd := exec.Command(c.git(), gitArgs...)
	cmd.Stderr = os.Stderr

	buf, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
				if waitStatus.ExitStatus() == 1 {
					return "", notFound(
						fmt.Sprintf("config not found. args: %q", strings.Join(gitArgs, " ")))
				}
			}
		}
		return "", err
	}
	return strings.TrimRight(string(buf), "\x00"), nil
}

// Get a value
func (c *Config) Get(args ...string) (string, error) {
	return c.Do(append([]string{"--get"}, args...)...)
}

// GetAll values
func (c *Config) GetAll(args ...string) ([]string, error) {
	val, err := c.Do(append([]string{"--get-all"}, args...)...)
	if err != nil {
		return nil, err
	}
	// No results found, return an empty slice
	if val == "" {
		return nil, nil
	}
	return strings.Split(val, "\x00"), nil
}

// Bool gets a value as bool
func (c *Config) Bool(key string) (bool, error) {
	val, err := c.Get("--bool", key)
	if err != nil {
		return false, err
	}
	return val == "true", nil
}

// Path gets a value as path
func (c *Config) Path(key string) (string, error) {
	return c.Get("--path", key)
}

// PathAll get all values as paths
func (c *Config) PathAll(key string) ([]string, error) {
	return c.GetAll("--path", key)
}

// Int get a value as int
func (c *Config) Int(key string) (int, error) {
	val, err := c.Get("--int", key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

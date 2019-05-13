package gitconfig

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Config struct {
	System, Global, Local bool
	File                  string
}

func (c *Config) Do(args ...string) (string, error) {
	gitArgs := append([]string{"config", "--null"})
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
	cmd := exec.Command("git", gitArgs...)
	cmd.Stderr = os.Stderr

	buf, err := cmd.Output()
	if exitError, ok := err.(*exec.ExitError); ok {
		if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
			if waitStatus.ExitStatus() == 1 {
				return "", notFound(
					fmt.Sprintf("config not found. args: %q", strings.Join(gitArgs, " ")))
			}
		}
		return "", err
	}
	return strings.TrimRight(string(buf), "\000"), nil
}

func (c *Config) Get(args ...string) (string, error) {
	return c.Do(append([]string{"--get"}, args...)...)
}

func (c *Config) GetAll(args ...string) ([]string, error) {
	val, err := c.Do(append([]string{"--get-all"}, args...)...)
	if err != nil {
		return nil, err
	}
	// No results found, return an empty slice
	if val == "" {
		return nil, nil
	}
	return strings.Split(val, "\000"), nil
}

func (c *Config) Bool(key string) (bool, error) {
	val, err := c.Get("--type=bool", key)
	if err != nil {
		return false, err
	}
	return val == "true", nil
}

func (c *Config) Path(key string) (string, error) {
	return c.Get("--type=path", key)
}

func (c *Config) PathAll(key string) ([]string, error) {
	return c.GetAll("--type=path", key)
}

func (c *Config) Int(key string) (int, error) {
	val, err := c.Get("--type=int", key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

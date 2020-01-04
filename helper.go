package gitconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// WithConfig is test helper to replace gitconfig temporarily
func WithConfig(tb testing.TB, configContent string) func() {
	tmpdir, err := ioutil.TempDir("", "gitconfig-test")
	if err != nil {
		if tb != nil {
			tb.Fatal(err)
		}
		panic(err)
	}

	tmpGitConfig := filepath.Join(tmpdir, "gitconfig")
	ioutil.WriteFile(tmpGitConfig, []byte(configContent), 0644)

	prevGitconfigEnv, ok := os.LookupEnv("GIT_CONFIG")
	os.Setenv("GIT_CONFIG", tmpGitConfig)

	return func() {
		if ok {
			os.Setenv("GIT_CONFIG", prevGitconfigEnv)
		} else {
			os.Unsetenv("GIT_CONFIG")
		}
		os.RemoveAll(tmpdir)
	}
}

package gitconfig

import "testing"

func TestConfigAll(t *testing.T) {
	dummyKey := "ghq.non.existent.key"
	confs, err := GetAll(dummyKey)
	if !IsNotFound(err) {
		t.Errorf("error should be notFounder, but: %s", err)
	}
	if len(confs) > 0 {
		t.Errorf("ConfigAll(%q) = %v; want %v", dummyKey, confs, nil)
	}
}

func TestConfigURL(t *testing.T) {
	defer WithConfig(t, `[ghq "https://ghe.example.com/"]
vcs = github
[ghq "https://ghe.example.com/hg/"]
vcs = hg
`)()

	testCases := []struct {
		name   string
		config []string
		expect string
	}{{
		name:   "github",
		config: []string{"--get-urlmatch", "ghq.vcs", "https://ghe.example.com/foo/bar"},
		expect: "github",
	}, {
		name:   "hg",
		config: []string{"--get-urlmatch", "ghq.vcs", "https://ghe.example.com/hg/repo"},
		expect: "hg",
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value, err := Do(tc.config...)
			if err != nil {
				t.Errorf("error should be nil but: %s", err)
			}
			if value != tc.expect {
				t.Errorf("got: %s, expect: %s", value, tc.expect)
			}
		})
	}
}

func TestConfigDo_nogit(t *testing.T) {
	c := &Config{GitPath: "dummy"}

	_, err := c.Do("hoge")
	if err == nil {
		t.Errorf("error shouldn't be nil")
	}
}

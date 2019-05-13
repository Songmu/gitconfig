package gitconfig

import (
	"testing"
)

func TestGetGHUserFromGH(t *testing.T) {
	user, err := getGHUserFromGHAPI("api.github.com", "y.songmu@gmail.com", "")
	if err != nil {
		t.Errorf("error should be nil, but: %s", err)
	}
	if user != "Songmu" {
		t.Errorf("getGHUserFromGHAPI() = %q, expect: Songmu", user)
	}

	user, err = getGHUserFromGHCommit("api.github.com", "y.songmu@gmail.com", "")
	if err != nil {
		t.Errorf("error should be nil, but: %s", err)
	}
	if user != "Songmu" {
		t.Errorf("getGHUserFromGHCommit() = %q, expect: Songmu", user)
	}
}

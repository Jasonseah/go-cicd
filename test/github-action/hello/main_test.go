package hello

import (
	"cicdtest/pkg/github-action/hello"
	"testing"
)

func TestGreetsGitHub(t *testing.T) {
	result := hello.Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want ello GitHub Actions. Dev.to is awesome", result)
	}
}

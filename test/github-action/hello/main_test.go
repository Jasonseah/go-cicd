package hello

import (
	"cicdtest/pkg/github-action/hello"
	"testing"
)

func TestGreetsGitHub(t *testing.T) {
	result := hello.Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; want Hello GitHub actions", result)
	}
}

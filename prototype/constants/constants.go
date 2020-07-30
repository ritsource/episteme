package constants

import (
	"os/exec"
	"path"
	"strings"
)

var RepositoryRoot string = func() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return strings.Split(string(out), "\n")[0]
}()

var DEFAULT_DATA_INPUT_CONFIG_FILEPATH = path.Join(RepositoryRoot, "prototype/config.yaml")

var DEFAULT_DATA_OUTPUT_DIR = path.Join(RepositoryRoot, "prototype/.data/dst/prod")
var DEFAULT_POSTS_DATA_OUTPUT_FILEPATH = path.Join(RepositoryRoot, "prototype/.data/dst/prod/posts")
var DEFAULT_PINNED_WEBSITES_DATA_OUTPUT_FILEPATH = path.Join(RepositoryRoot, "prototype/.data/dst/prod/pinned_websites")

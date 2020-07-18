package constants

import (
	"os/exec"
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

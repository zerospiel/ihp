// Package git is for git subcommand.
package git

import (
	"fmt"
	"os/exec"
)

// Setup initializes git repo if it has no one.
//
// Then sets the given email and username as a local git config.
func Setup(email, uname string) error {
	if len(email) == 0 || len(uname) == 0 {
		return fmt.Errorf("both email and username should be setted")
	}

	args := []string{"git", "rev-parse", "--is-inside-work-tree"}
	if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
		if err = exec.Command(args[0], "init").Run(); err != nil {
			return fmt.Errorf("failed to init git repository: %w", err)
		}
	}

	if err := exec.Command(args[0], "config", "--local", "--add", "user.email", email).Run(); err != nil {
		return fmt.Errorf("failed to add email to git local config: %w", err)
	}
	if err := exec.Command(args[0], "config", "--local", "--add", "user.name", uname).Run(); err != nil {
		return fmt.Errorf("failed to add username to git local config: %w", err)
	}
	return nil
}

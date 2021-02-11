package services

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func GitClone(repo, dest string) error {
	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout,
	})

	return err
}

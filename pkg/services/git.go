package services

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type Git struct {
	Repo string
	Tag  string
}

func GitClone(repo, dest string) error {
	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL: fmt.Sprintf("https://%s", repo),
	})

	return err

	// if err != nil {
	// 	return err
	// }

	// w, err := r.Worktree()

	// if err != nil {
	// 	return err
	// }

	// return w.Checkout(&git.CheckoutOptions{
	// 	Branch: "v0.1.0",
	// })
}

package lib

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
	gitPlumbing "gopkg.in/src-d/go-git.v4/plumbing"
)

func getRefName(_refName string) gitPlumbing.ReferenceName {
	refName := gitPlumbing.ReferenceName(_refName)

	if _refName == "" {
		refName = gitPlumbing.ReferenceName("HEAD")
	}

	return refName
}

// Clone : Clones a given repository into targetDir
func Clone(url string, targetDir string, _refName string, _progress bool) (r *git.Repository, err error) {
	options := &git.CloneOptions{
		URL:           url,
		Depth:         1,
		ReferenceName: getRefName(_refName),
		SingleBranch:  true,
		Progress:      os.Stdout,
	}

	if _progress {
		options.Progress = os.Stdout
	}

	return git.PlainClone(
		targetDir,
		false,
		options,
	)
}

// Pull : Pulls a given repository
func Pull(targetDir string, _refName string, _progress bool) (err error) {
	r, err := git.PlainOpen(targetDir)

	if err != nil {
		panic(err)
	}

	options := &git.PullOptions{
		ReferenceName: getRefName(_refName),
	}

	if _progress {
		options.Progress = os.Stdout
	}

	return r.Pull(options)
}

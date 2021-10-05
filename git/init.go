package git

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func (r *gitRepository) Init(path string, force bool) error {
	r.worktree = path
	r.gitDir = filepath.Join(r.worktree, ".git")
	fileinfo, err := os.Stat(r.gitDir)
	if err != nil && !force {
		return err
	}
	if fileinfo != nil && (!fileinfo.IsDir() || !force) {
		return fmt.Errorf("file path is not a directory: %s", path)
	}

	ini, err := ini.Load(r.repoPath("config"))
	if err != nil && !force {
		return fmt.Errorf("could not read git config file")
	}

	if !force {
		gitVersion := ini.Section("").Key("repositoryformatversion").String()
		if gitVersion != "0" {
			return fmt.Errorf("incorrect git version in config file: %s", gitVersion)
		}
	}
	err = r.create()
	return err
}

func (r *gitRepository) create() error {
	fileinfo, err := os.Stat(r.worktree)
	if err == nil {
		if !fileinfo.IsDir() {
			return fmt.Errorf("%s is not a directory", r.worktree)
		}
		if !isEmptyDir(r.worktree) {
			return fmt.Errorf("%s is not empty", r.worktree)
		}
	} else {
		os.MkdirAll(r.worktree, 0770)
	}

	r.repoDir("branches")
	r.repoDir("objects")
	r.repoDir("refs/tags")
	r.repoDir("refs/heads")

	desc := []byte("Unnamed repository; edit this file 'description' to name the repository.\n")
	err = os.WriteFile(r.repoPath("description"), desc, 0644)
	if err != nil {
		return err
	}

	head := []byte("ref: refs/heads/master\n")
	err = os.WriteFile(r.repoPath("HEAD"), head, 0644)
	if err != nil {
		return err
	}

	cfg := ini.Empty()
	if err != nil {
		return err
	}

	coreSection, err := cfg.NewSection("core")
	if err != nil {
		return err
	}
	coreSection.NewKey("repositoryformatversion", "0")
	coreSection.NewKey("filemode", "true")
	coreSection.NewKey("bare", "false")
	f, err := os.Create(r.repoPath("config"))
	if err != nil {
		return err
	}
	cfg.WriteTo(f)
	return nil
}

package git

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type gitRepository struct {
	worktree string
	gitDir   string
	conf     string
}

func New() *gitRepository {
	return &gitRepository{}
}

func (r *gitRepository) Init(path string, force bool) (*gitRepository, error) {
	r.worktree = path
	r.gitDir = filepath.Join(r.worktree, ".git")
	fileinfo, err := os.Stat(r.gitDir)
	if err != nil && !force {
		return &gitRepository{}, err
	}
	if fileinfo != nil && (!fileinfo.IsDir() || !force) {
		return &gitRepository{}, fmt.Errorf("file path is not a directory: %s", path)
	}

	ini, err := ini.Load(r.repoPath("config"))
	if err != nil && !force {
		return &gitRepository{}, fmt.Errorf("could not read git config file")
	}

	if !force {
		gitVersion := ini.Section("").Key("repositoryformatversion").String()
		if gitVersion != "0" {
			return &gitRepository{}, fmt.Errorf("incorrect git version in config file: %s", gitVersion)
		}
	}
	err = r.create()
	return r, err
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
	f, _ := os.Create(r.repoPath("config"))
	cfg.WriteTo(f)
	return nil
}

func (r *gitRepository) repoDir(path string) (string, error) {
	path = r.repoPath(path)
	fileinfo, err := os.Stat(path)
	if err == nil { // file/dir exists
		if !fileinfo.IsDir() {
			return "", fmt.Errorf("%s is not a directory", r.worktree)
		}
		return path, nil
	}

	os.MkdirAll(path, 0770)
	return path, nil
}

func (r *gitRepository) repoPath(path string) string {
	return filepath.Join(r.gitDir, path)
}

func isEmptyDir(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdir(1)
	return err == io.EOF
}

package git

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Repository struct {
	worktree string
	gitDir   string
}

func NewExistingRepo() *Repository {
	gitpath, err := FindGitDir(".")
	if err != nil {
		log.Fatalf("err finding git dir: %v", err)
	}
	return NewRepo(strings.TrimSuffix(gitpath, ".git"))
}

func NewRepo(path string) *Repository {
	return &Repository{
		gitDir:   fmt.Sprintf("%s.git", path),
		worktree: path,
	}
}

func (r *Repository) Gitdir() string {
	return r.gitDir
}

func (r *Repository) repoDir(path string, create bool) (*string, error) {
	path = r.repoPath(path)
	fileinfo, err := os.Stat(path)
	if err == nil { // file/dir exists
		if !fileinfo.IsDir() {
			return nil, fmt.Errorf("%s is not a directory", r.worktree)
		}
		return &path, nil
	}

	if create {
		err = os.MkdirAll(path, 0770)
		if err != nil {
			return nil, err
		}
	}
	return &path, nil
}

func (r *Repository) repoPath(path string) string {
	return filepath.Join(r.gitDir, path)
}

func (r *Repository) RepoFile(path string, create bool) string {
	_, err := r.repoDir(filepath.Dir(path), create)
	if err != nil {
		return ""
	}
	return path
}

func isEmptyDir(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	// try to read one file from dir. If the dir is empty it will return a io.EOF error.
	_, err = f.Readdir(1)
	return err == io.EOF
}

func FindGitDir(path string) (string, error) {
	if p, _ := filepath.Abs(path); p == "/" {
		return "", fmt.Errorf("no git repo found")
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, f := range files {
		if f.Name() == ".git" && f.IsDir() {
			return strings.Join([]string{path, f.Name()}, "/"), nil
		}
	}

	path, err = FindGitDir(filepath.Join(path, ".."))
	if err != nil {
		return "", err
	}

	return path, nil
}

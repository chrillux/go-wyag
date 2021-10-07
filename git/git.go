package git

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type gitRepository struct {
	worktree string
	gitDir   string
}

func New() *gitRepository {
	return &gitRepository{}
}

func (r *gitRepository) repoDir(path string, create bool) (*string, error) {
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

func (r *gitRepository) repoPath(path string) string {
	return filepath.Join(r.gitDir, path)
}

func (r *gitRepository) repoFile(path string, create bool) string {
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

	_, err = f.Readdir(1)
	return err == io.EOF
}

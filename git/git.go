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

func (r *gitRepository) repoDir(path string) (*string, error) {
	path = r.repoPath(path)
	fileinfo, err := os.Stat(path)
	if err == nil { // file/dir exists
		if !fileinfo.IsDir() {
			return nil, fmt.Errorf("%s is not a directory", r.worktree)
		}
		return &path, nil
	}

	err = os.MkdirAll(path, 0770)
	if err != nil {
		return nil, err
	}
	return &path, nil
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

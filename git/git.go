package git

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
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

func (r *Repository) Init(path string, force bool) error {
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

func (r *Repository) create() error {
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

	r.repoDir("branches", true)
	r.repoDir("objects", true)
	r.repoDir("refs/tags", true)
	r.repoDir("refs/heads", true)

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

func (r *Repository) RefResolve(ref string) string {
	refpath := r.RepoFile(ref, false)
	data, err := os.ReadFile(refpath)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}
	sdata := string(data)
	if strings.HasPrefix(sdata, "ref: ") {
		return r.RefResolve(strings.TrimPrefix(sdata, "ref: "))
	}
	return sdata
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

package storage

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Storage interface {
	FetchFiles(ctx context.Context) ([][]byte, error)
}

type GitStorage struct {
	repoURL string

	cloneFunc func(url string) (*git.Repository, error)
}

// NewGitStorage creates a new GitStorage with the given repoURL.
func NewGitStorage(repoURL, username, password string) *GitStorage {
	return &GitStorage{
		repoURL: repoURL,
		cloneFunc: func(url string) (*git.Repository, error) {
			return git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
				URL: url,
				Auth: &http.BasicAuth{
					Username: username,
					Password: password,
				},
			})
		},
	}
}

func (gs *GitStorage) FetchFiles(ctx context.Context) ([][]byte, error) {
	r, err := gs.cloneFunc(gs.repoURL)

	if err != nil {
		return nil, fmt.Errorf("failed to clone: %w", err)
	}

	ref, err := r.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD reference: %w", err)
	}

	// Get the commit object for the HEAD reference
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get commit object: %w", err)
	}

	// Get the tree associated with the commit
	tree, err := commit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed to get tree: %w", err)
	}

	// Filter and read .yml files in the tree

	files := [][]byte{}
	err = tree.Files().ForEach(func(f *object.File) error {
		ext := filepath.Ext(f.Name)
		if ext == ".yml" || ext == ".yaml"{			

			// Read the file content
			reader, err := f.Reader()
			if err != nil {
				return fmt.Errorf("failed to read file content: %w", err)
			}
			defer reader.Close()

			content, err := io.ReadAll(reader)
			if err != nil {
				return fmt.Errorf("failed to read file content: %w", err)
			}

			files = append(files, content)

		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to process files: %w", err)
	}

	return files, nil
}

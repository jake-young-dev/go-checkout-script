package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/afero"
)

const (
	CLONE_TIMEOUT = time.Second * 10 //timeout for git clone
)

func main() {
	workingdir := os.Getenv("GH_WORKSPACE")
	fmt.Printf("using workspace %s\n", workingdir)

	fs := afero.NewOsFs()
	fmt.Println("clearing files from workspace")
	err := deleteEverything(fs, workingdir)
	if err != nil {
		panic(err)
	}
	fmt.Println("cleared!")

	ctx, cancel := context.WithTimeout(context.Background(), CLONE_TIMEOUT)
	defer cancel()

	var opts *git.CloneOptions
	if _, ok := os.LookupEnv("GH_REPO_LINK"); ok {
		repo := os.Getenv("GH_REPO_LINK")
		fmt.Printf("pulling repo from input: %s\n", repo)
		opts = &git.CloneOptions{
			URL: repo,
		}
	} else {
		repo := os.Getenv("GH_DEFAULT_REPO")
		fmt.Printf("pulling current code repo: %s\n", repo)
		opts = &git.CloneOptions{
			URL: repo,
		}
	}
	fmt.Println("cloning")
	_, err = git.PlainCloneContext(ctx, workingdir, false, opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("cloned!")

}

func deleteEverything(fs afero.Fs, dir string) error {
	return afero.Walk(fs, dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if path == dir {
			return nil
		}

		if info.IsDir() {
			c, err := afero.Exists(fs, path)
			if err != nil {
				return err
			}
			if !c {
				return nil
			}
			err = fs.RemoveAll(path)
			if err != nil {
				return err
			}
		} else {
			c, err := afero.Exists(fs, path)
			if err != nil {
				return err
			}
			if !c {
				return nil
			}
			err = fs.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

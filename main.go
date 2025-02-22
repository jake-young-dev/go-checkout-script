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
	//setting github working directory
	workingdir := os.Getenv("GH_WORKSPACE")
	fmt.Printf("using workspace %s\n", workingdir)

	fs := afero.NewOsFs()
	fmt.Printf("cleaning %s\n", workingdir)
	err := deleteEverything(fs, workingdir)
	if err != nil {
		panic(err)
	}
	fmt.Println("cleaned")

	ctx, cancel := context.WithTimeout(context.Background(), CLONE_TIMEOUT)
	defer cancel()

	var opts *git.CloneOptions
	repoInput := os.Getenv("GH_REPO_LINK")
	if repoInput != "" {
		fmt.Println("pulling input repo")
		opts = &git.CloneOptions{
			URL: repoInput,
		}
	} else {
		fmt.Println("pulling current repo")
		opts = &git.CloneOptions{
			URL: os.Getenv("GH_DEFAULT_REPO"),
		}
	}

	fmt.Printf("cloning %s into %s\n", opts.URL, workingdir)
	_, err = git.PlainCloneContext(ctx, workingdir, false, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cloned %s\n", opts.URL)

}

// deletes all files in current working directory so we can clone a new repository into it
func deleteEverything(fs afero.Fs, dir string) error {
	return afero.Walk(fs, dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil //don't stop on "errors"
		}
		if path == dir {
			return nil //don't delete working directory
		}

		if info.IsDir() {
			c, err := afero.Exists(fs, path) //have to check existance to prevent errors
			if err != nil {
				return err
			}
			if !c {
				return nil
			}
			err = fs.RemoveAll(path) //remove folder and children
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
			err = fs.Remove(path) //remove file
			if err != nil {
				return err
			}
		}
		return nil
	})
}

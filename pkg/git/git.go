package git

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func Commit() {
	r, err := git.PlainOpen("/tmp/foo")
	if err != nil {
		panic(err)
	}

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	filename := filepath.Join("/tmp/foo", "foo-bar-file")
	err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)

	_, err = w.Add("foo-bar-file")
	status, err := w.Status()
	fmt.Println(status)

	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	obj, err := r.CommitObject(commit)
	fmt.Println(obj)
}

var (
	errReachedToCommit = errors.New("reached to commit")
)

func Log() {
	r, _ := git.PlainOpen("/tmp/foo")

	tIter, _ := r.TagObjects()
	_ = tIter.ForEach(func(tag *object.Tag) error {
		fmt.Println(tag.Name)
		return nil
	})

	tIter1, _ := r.Tags()
	_ = tIter1.ForEach(func(reference *plumbing.Reference) error {
		fmt.Println(reference.Target())
		return nil
	})

	//fmt.Println(commits)
	/*
		cIter, err := r.Log(&git.LogOptions{From: plumbing.NewHash("447e58cb2f2f51a89293f9c40566da5dc224cbf6")})
		err = cIter.ForEach(func(c *object.Commit) error {
			fmt.Println(c.Message)

			return nil
		})
	*/
}

func Config() {
	r, err := git.PlainOpen("/tmp/foo")
	if err != nil {
		panic(err)
	}

	c, err := r.ConfigScoped(config.GlobalScope)
	if err != nil {
		panic(err)
	}

	fmt.Println(c.User.Name)
	fmt.Println(c.User.Email)
}

func Main() {
	Log()
}

package comfyUI

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"os/exec"
)

func (c *ComfyUI) GitHeadCommit(comfyUIPath string) *object.Commit {
	repo, err := git.PlainOpen(comfyUIPath)
	if err != nil {
		return nil
	}
	head, err := repo.Head()
	if err != nil {
		return nil
	}
	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		return nil
	}
	return commit
}

func (c *ComfyUI) GitLatestCommit(comfyUIPath string) error {
	repo, err := git.PlainOpen(comfyUIPath)
	if err != nil {
		return err
	}
	head, err := repo.Head()
	if err != nil {
		return err
	}
	remote, err := repo.Remote("origin")

	go func() {
		err := remote.Fetch(&git.FetchOptions{})
		if err != nil {
		}

		reference, err := repo.Reference(plumbing.NewRemoteReferenceName("origin", head.Name().Short()), true)
		if err != nil {
		}
		reference.Hash()

		commit, err := repo.CommitObject(reference.Hash())
		if err != nil {
		}
		runtime.EventsEmit(c.ctx, "GitLatestCommit", commit)
	}()

	return err
}

func (c *ComfyUI) GitPull(comfyUIPath string) string {
	cmd := exec.Command("git", "pull")
	cmd.Dir = comfyUIPath
	output, err := cmd.CombinedOutput()
	log.Println(err)
	return string(output)
}

func (c *ComfyUI) GitStatus(comfyUIPath string) []*object.Commit {
	behindCommits := make([]*object.Commit, 0)

	repo, err := git.PlainOpen(comfyUIPath)
	if err != nil {
		fmt.Printf("无法打开仓库: %v\n", err)
		return nil
	}

	// 获取当前分支
	head, err := repo.Head()
	if err != nil {
		fmt.Printf("无法获取 HEAD: %v\n", err)
		return nil
	}

	// 获取远程分支引用
	remoteRef, err := repo.Reference(plumbing.NewRemoteReferenceName("origin", head.Name().Short()), true)
	if err != nil {
		fmt.Printf("无法获取远程引用: %v\n", err)
		return nil
	}

	// 计算落后的提交数
	localHash := head.Hash()
	remoteHash := remoteRef.Hash()

	divergence, err := repo.Log(&git.LogOptions{From: remoteHash})
	if err != nil {
		fmt.Printf("无法计算分支落后情况: %v\n", err)
		return nil
	}

	err = divergence.ForEach(func(c *object.Commit) error {
		if c.Hash == localHash {
			return storer.ErrStop
		}
		behindCommits = append(behindCommits, c)
		return nil
	})
	return behindCommits
}

package components

import (
	"fmt"
	"os"
	"strings"

	"github.com/sihuayin/godist/models"
	"github.com/sihuayin/godist/pkg/ssh"
)

type GitComponent struct {
	Project *models.Project
}

func NewGitComponent(project *models.Project) *GitComponent {
	return &GitComponent{
		Project: project,
	}
}

func (git *GitComponent) GetBranchList() ([]map[string]string, error) {
	history := []map[string]string{}

	destination := git.Project.GetDeployFromDir()
	fmt.Println("destination -> ", destination)
	git.UpdateRepo("master", destination)
	return history, nil
}

func (git *GitComponent) UpdateRepo(branch, gitDir string) error {
	if gitDir == "" {
		gitDir = git.Project.GetDeployFromDir()
	}

	if branch == "" {
		branch = "master"
	}
	dotGit := strings.TrimRight(gitDir, "/") + "/.git"
	if _, err := os.Stat(dotGit); err != nil {
		if os.IsNotExist(err) {
			cmds := []string{}
			cmds = append(cmds, fmt.Sprintf("mkdir -p %s ", gitDir))
			cmds = append(cmds, fmt.Sprintf("cd %s ", gitDir))
			cmds = append(cmds, fmt.Sprintf("/usr/bin/env git clone -q %s", git.Project.RepoUrl))
			cmds = append(cmds, fmt.Sprintf("/usr/bin/env git checkout -q %s", branch))
			cmd := strings.Join(cmds, " && ")
			err := git.runLocalCommand(cmd)
			return err
		}
	}

	return nil
}

func (git *GitComponent) runLocalCommand(command string) error {
	s, err := ssh.CommandLocal(command, 3600)
	fmt.Println(command)
	fmt.Println(s, err)
	return err
}

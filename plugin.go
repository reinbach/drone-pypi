package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

type (
	Build struct {
		Home string
		Workspace string
	}

	Config struct {
		Repo string
		Username string
		Password string
		Distributions []string
	}

	Plugin struct {
		Build Build
		Config Config
	}
)

func (p Plugin) CreateConfig() error {
	f, err := os.Create(path.Join(p.Build.Home, ".pypirc"))
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	err = p.WriteConfig(buf)
	if err != nil {
		return err
	}
	buf.Flush()
	return nil
}

func (p Plugin) WriteConfig(w io.Writer) error {
	_, err := io.WriteString(
		w,
		fmt.Sprintf(
			`[distutils]
index-servers =
    pypi

[pypi]
repository: %s
username: %s
password: %s
`,
			p.Config.Repo,
			p.Config.Username,
			p.Config.Password,
		),
	)

	return err
}

func (p Plugin) UploadDist() error {
	distributions := p.Config.Distributions

	args := []string{"setup.py"}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	args = append(args, "upload")
	args = append(args, "-r")
	args = append(args, "pypi")

	cmd := exec.Command("python", args...)
	cmd.Dir = p.Build.Workspace
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("$", strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func (p Plugin) Exec() error {
	err := p.CreateConfig()
	if err != nil {
		fmt.Printf("Error: Failed to create .pypirc file. %s\n", err)
		return err
	}
	err = p.UploadDist()
	if err != nil {
		fmt.Printf("Error: Failed to create/upload distribution. %s\n", err)
		return err
	}

	return nil
}

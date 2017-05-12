package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func CreateConfig() error {
	f, err := os.Create(path.Join(os.Getenv("HOME"), ".pypirc"))
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	err = WriteConfig(buf)
	if err != nil {
		return err
	}
	buf.Flush()
	return nil
}

func WriteConfig(w io.Writer) error {
	repo := os.Getenv("PLUGIN_REPOSITORY")
	if repo == "" {
		repo = "https://pypi.python.org/pypi"
	}
	username := os.Getenv("PLUGIN_USERNAME")
	password := os.Getenv("PLUGIN_PASSWORD")
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
			repo,
			username,
			password,
		),
	)

	return err
}

func UploadDist() error {
	dists := os.Getenv("PLUGIN_DISTRIBUTIONS")
	distributions := strings.Split(dists, ",")

	args := []string{"setup.py"}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	args = append(args, "upload")
	args = append(args, "-r")
	args = append(args, "pypi")

	cmd := exec.Command("python", args...)
	cmd.Dir = os.Getenv("DRONE_WORKSPACE")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("$", strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func main() {
	err := CreateConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = UploadDist()
	if err != nil {
		log.Fatal(err)
	}
}

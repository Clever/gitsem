package main

import (
	"bytes"
	"flag"
	"fmt"
	"gopkg.in/blang/semver.v1"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getCurrentVersion(path string) (*semver.Version, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &semver.Version{}, nil
	}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return semver.New(string(contents))
}

func getRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	result := &bytes.Buffer{}
	cmd.Stdout = result
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(result.String()), nil
}

func isRepoClean() (bool, error) {
	cmd := exec.Command("git", "status", "-s")
	result := &bytes.Buffer{}
	cmd.Stdout = result
	if err := cmd.Run(); err != nil {
		return false, err
	}
	return result.String() == "", nil
}

func addFile(path string) error {
	return exec.Command("git", "add", path).Run()
}

func commit(message string) error {
	return exec.Command("git", "commit", "-m", message).Run()
}

func commitMessage(message, version string) string {
	if message == "" {
		return version
	} else if strings.Contains(message, "%s") {
		return fmt.Sprintf(message, version)
	} else {
		return message
	}
}

func tag(version string) error {
	return exec.Command("git", "tag", version).Run()
}

const versionFileName = "VERSION"

func main() {
	message := flag.String("m", "", "commit message for version commit")
	help := flag.Bool("h", false, "print usage and exit")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if clean, err := isRepoClean(); err != nil {
		log.Fatal(err)
	} else if !clean {
		log.Fatal("repo isn't clean")
	}

	root, err := getRepoRoot()
	if err != nil {
		log.Fatal(err)
	}
	versionFile := filepath.Join(root, versionFileName)
	version, err := getCurrentVersion(versionFile)
	if err != nil {
		log.Fatal(err)
	}
	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "missing version argument\n\n")
		flag.Usage()
		os.Exit(1)
	}

	newVersion := flag.Args()[0]
	switch newVersion {
	case "patch":
		version.Patch++
	case "minor":
		version.Minor++
	case "major":
		version.Major++
	default:
		if version, err = semver.New(newVersion); err != nil {
			log.Fatal(err)
		}
	}

	if err := ioutil.WriteFile(versionFile, []byte(version.String()), 0666); err != nil {
		log.Fatal(err)
	}
	if err := addFile(versionFile); err != nil {
		log.Fatal(err)
	}
	versionString := "v" + version.String()
	*message = commitMessage(*message, versionString)
	if err := commit(*message); err != nil {
		log.Fatal(err)
	}
	if err := tag(versionString); err != nil {
		log.Fatal(err)
	}
}

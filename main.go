package main

import (
	"flag"
	"fmt"
	"gopkg.in/blang/semver.v1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

const versionFileName = "VERSION"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: [options] version\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "version can be one of: newversion | patch | minor | major\n\n")
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
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

	root, err := repoRoot()
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
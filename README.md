# gitsem

A command line utility for managing semantically versioned (semver) git tags.

Run this in a git repository to bump the version and write the new data back to the VERSION file.
It will also create a version commit and (optional) tag, and fail if the repo is not clean.

## Installation

```shell
$ go get github.com/Clever/gitsem
```

## Example
```shell
$ gitsem patch
$ gitsem -m "Upgrade to %s for reasons" patch
$ gitsem minor
```

## Usage

```shell
gitsem [options] version
```

`version` can be one of: `newversion | patch | minor | major`

The version argument should be a valid semver string, or a field of a semver string (one of "patch", "minor", or "major").
In the second case, the existing version will be incremented by 1 in the specified field.

### Options
  - `m=%s` specifies a commit message to use when bumping the version. If %s appears, it will be replaced with the new version number.
  - `tag=true` whether or not to create a tag at the version commit

## Changing Dependencies

### New Packages

When adding a new package, you can simply use `make vendor` to update your imports.
This should bring in the new dependency that was previously undeclared.
The change should be reflected in [Godeps.json](Godeps/Godeps.json) as well as [vendor/](vendor/).

### Existing Packages

First ensure that you have your desired version of the package checked out in your `$GOPATH`.

When to change the version of an existing package, you will need to use the godep tool.
You must specify the package with the `update` command, if you use multiple subpackages of a repo you will need to specify all of them.
So if you use package github.com/Clever/foo/a and github.com/Clever/foo/b, you will need to specify both a and b, not just foo.

```
# depending on github.com/Clever/foo
godep update github.com/Clever/foo

# depending on github.com/Clever/foo/a and github.com/Clever/foo/b
godep update github.com/Clever/foo/a github.com/Clever/foo/b
```


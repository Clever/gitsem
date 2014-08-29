# gitsem

a command line utility for managing semantically versioned (semver) git tags

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

Run this in a git repository to bump the version and write the new data back to the VERSION file.

The newversion argument should be a valid semver string, or a valid second argument to semver.inc (one of "patch", "minor", or "major").
In the second case, the existing version will be incremented by 1 in the specified field.

It will also create a version commit and tag, and fail if the repo is not clean.

If supplied with --message (shorthand: -m) config option, gitsem will use it as a commit message when creating a version commit.
If the message config contains %s then that will be replaced with the resulting version number.

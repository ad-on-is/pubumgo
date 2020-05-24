# PuBumGo

Made for Flutter with ❤ in Go

## What is it?

A simple CLI tool to bump the version in the pubspec.yaml file of your Flutter project.

Can be used in git hooks.

## Usage

`pubumgo arg [build]`

## Valid args

- `major`: bumps the version from `1.2.3+4` to `2.0.0`
- `minor`: bumps the version from `1.2.3+4` to `1.3.0`
- `patch`: bumps the version from `1.2.3+4` to `1.2.4`
- `build`: bumps the version from `1.2.3+4` to `1.2.3+5`

Use `-b` additionally to preserve the build-number

`pubumgo release -b` will bumpt the version from `1.2.3+4` to `2.0.0+4`

## Use as a git-hook, example `project-root/.git/hooks/pre-commit`

```
#!/bin/bash

pubumgo patch
git add pubspec.yaml
```

Then in your `project-root` just invoke the git commands as usual, `git add`, and `git commit`, and your pubspec version will automatically be bumped.

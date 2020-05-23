# PuBumGo

Made for Flutter with ❤ in Go

## What is it?

A simple CLI tool to bump the version in the pubspec.yaml file of your Flutter project.

Can be used in git hooks.

## Usage

`pubumgo arg [build]`

## Valid args

- `release`: bumps the release from `1.2.3+4` to `2.0.0`
- `major`: bumps the version from `1.2.3+4` to `1.3.0`
- `minor`: bumps the version from `1.2.3+4` to `1.2.4`
- `build`: bumps the version from `1.2.3+4` to `1.2.3+5`

Use `build` additionally to preserve the build-number

`pubumgo release build` will bumpt the version from `1.2.3+4` to `2.0.0+4`

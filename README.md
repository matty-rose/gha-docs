# gha-docs

[![test](https://img.shields.io/github/workflow/status/matty-rose/gha-docs/Test%20GHA%20Docs)](https://github.com/matty-rose/gha-docs/actions/workflows/test.yaml)
[![coverage](https://img.shields.io/codecov/c/github/matty-rose/gha-docs)](https://codecov.io/gh/matty-rose/gha-docs)
[![Go Reference](https://pkg.go.dev/badge/github.com/matty-rose/gha-docs.svg)](https://pkg.go.dev/github.com/matty-rose/gha-docs)
[![Go Report Card](https://goreportcard.com/badge/github.com/matty-rose/gha-docs)](https://goreportcard.com/report/github.com/matty-rose/gha-docs)
[![license](https://img.shields.io/github/license/matty-rose/gha-docs)](https://github.com/matty-rose/gha-docs/blob/main/LICENSE)
[![release](https://img.shields.io/github/v/release/matty-rose/gha-docs)](https://github.com/matty-rose/gha-docs/releases)

## What is it?

A program to generate documentation for GitHub [composite actions](https://docs.github.com/en/actions/creating-actions/creating-a-composite-action).

## Installation

### Stable Release Binaries
Substitute `darwin` for `linux` if installing on MacOS.

```bash
export GHA_DOCS_VERSION=v0.11.0
export GHA_DOCS_OS=linux
export GHA_DOCS_ARCH=amd64

curl -Lo ./gha-docs.tar.gz https://github.com/matty-rose/gha-docs/releases/download/$GHA_DOCS_VERSION/gha-docs-$GHA_DOCS_VERSION-$GHA_DOCS_OS-$GHA_DOCS_ARCH.tar.gz
tar -xzf gha-docs.tar.gz
chmod +x gha-docs && mv gha-docs /usr/local/bin
```

### From Source

Go `1.17+` is required.

```bash
go install github.com/matty-rose/gha-docs@v0.11.0
```

## Usage

Run `gha-docs` to display general usage.

### Generating a README
While only writing to stdout is supported:
```bash
gha-docs generate path/to/action.yaml > README.md
```

## License

MIT License - Copyright (c) 2021

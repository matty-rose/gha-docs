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
export GHA_DOCS_VERSION=v0.15.1
export GHA_DOCS_OS=linux
export GHA_DOCS_ARCH=amd64

curl -Lo ./gha-docs.tar.gz https://github.com/matty-rose/gha-docs/releases/download/$GHA_DOCS_VERSION/gha-docs-$GHA_DOCS_VERSION-$GHA_DOCS_OS-$GHA_DOCS_ARCH.tar.gz
tar -xzf gha-docs.tar.gz
chmod +x gha-docs && mv gha-docs /usr/local/bin
```

### From Source

Go `1.17+` is required.

```bash
go install github.com/matty-rose/gha-docs@v0.15.1
```

## Usage

Run `gha-docs` to display general usage.

To output to stdout:
```bash
gha-docs generate path/to/action.yaml
```

### Generating a README

This will overwrite any existing content in `README.md`.
```bash
gha-docs generate --output-file README.md path/to/action.yaml
```

To inject generated documentation into an existing file, adding the following markers in the file will cause the documentation to be generated between them, overwriting any content already existing between the markers, and preserving any content outside the markers.
```md
<!-- BEGIN GHA DOCS -->
<!-- END GHA DOCS -->
```

Then use the `-i/--inject` flag to inject the documentation into the file specified by `-o/--output-file` e.g.
```bash
gha-docs generate -i -o README.md path/to/action.yaml
```

## Future Improvements
- [ ] Add CI workflow to auto-commit find/replace version updates to README on release
- [ ] Parse config from `.gha-docs.yml`

## License

MIT License - Copyright (c) 2021

# gha-docs

## Installation
Substitute `darwin` for `linux` if installing on MacOS.

```bash
export GHA_DOCS_VERSION=v0.11.0
export GHA_DOCS_OS=linux
export GHA_DOCS_ARCH=amd64

curl -LO https://github.com/matty-rose/gha-docs/releases/download/$GHA_DOCS_VERSION/gha-docs-$GHA_DOCS_VERSION-$GHA_DOCS_OS-$GHA_DOCS_ARCH.tar.gz
tar -xzf gha-docs-$GHA_DOCS_VERSION-$GHA_DOCS_OS-$GHA_DOCS_ARCH.tar.gz
chmod +x gha-docs && mv gha-docs /usr/local/bin
```

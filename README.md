# Farcloser containers library

A very simple library providing basics for the containers ecosystem.

This contains modification of code, copies of code, or references to code from:
- github.com/containerd/nerdctl
- github.com/containerd/containerd
- github.com/moby/moby

and others, as indicated in the relevant files.

All copied and original code licensed under the Apache License.

## Dev

### Requirements

Review or call directly `./hack/dev-setup-linux.sh` or `./hack/dev-setup-macos.sh`

Then, `make install-dev-tools`

### Makefile

```bash
make fix
make lint
make test
```

### Local documentation

```bash
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite
open http://localhost:8080/go.farcloser.world/containers
```

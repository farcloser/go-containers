# Farcloser containers library

A very simple library providing basics for the containers ecosystem.

At this point, this is merely hiding away `opencontainers/*`.

## Dev

### Makefile

```bash
make lint
make lint-fix
make tidy
```

### Local documentation

```bash
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite
open http://localhost:8080/go.farcloser.world/containers
```

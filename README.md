gitconfig
=======

[![Test Status](https://github.com/Songmu/gitconfig/workflows/test/badge.svg?branch=main)][actions]
[![codecov.io](https://codecov.io/github/Songmu/gitconfig/coverage.svg?branch=main)][codecov]
[![MIT License](https://img.shields.io/github/license/Songmu/gitconfig)][license]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/Songmu/gitconfig)][PkgGoDev]

[actions]: https://github.com/Songmu/gitconfig/actions?workflow=test
[codecov]: https://codecov.io/github/Songmu/gitconfig?branch=main
[license]: https://github.com/Songmu/gitconfig/blob/main/LICENSE
[PkgGoDev]: https://pkg.go.dev/github.com/Songmu/gitconfig

gitconfig is a package to get configuration values from gitconfig.

## Synopsis

```go
val, err := gitconfig.Get("section.value")
if err != nil && !gitconfig.IsNotFound(err) {
    return err
}

// detect GitHub username from various informations
u, err := gitconfig.GitHubUser("")
```

## Description

## Installation

```console
% go get github.com/Songmu/gitconfig
```

## Author

[Songmu](https://github.com/Songmu)

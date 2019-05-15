gitconfig
=======

[![Build Status](https://travis-ci.org/Songmu/gitconfig.svg?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/Songmu/gitconfig/badge.svg?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/gitconfig?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/gitconfig
[coveralls]: https://coveralls.io/r/Songmu/gitconfig?branch=master
[license]: https://github.com/Songmu/gitconfig/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/gitconfig

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

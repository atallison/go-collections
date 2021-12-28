# Collection

[![Go Reference](https://pkg.go.dev/badge/github.com/hidetatz/collection.svg)](https://pkg.go.dev/github.com/hidetatz/collection)
[![.github/workflows/test.yml](https://github.com/hidetatz/collection/actions/workflows/test.yml/badge.svg)](https://github.com/hidetatz/collection/actions/workflows/test.yml)

Collection library using generics in Go

## Overview

This is a library to provide useful collection data structures and methods for Gophers.
This library uses generics feature which will get available in Go officially on 1.18.

## Usage

All of the data structures are **not** concurrent safe by default.
The synchronization to avoid data race is the caller's responsibility.

## Prerequisites

Go 1.18 or later.

If Go 1.18 has not been released but you want to try this package, you can use gotip command.

```go
$ go install golang.org/dl/gotip@latest
$ gotip download # latest commit
$ gotip version
go version devel go1.18-c2397905e0 Sat Nov 13 03:33:55 2021 +0000 darwin/arm64
```

You must use `gotip build` to build this library instead of `go build` .

## Documentation

As of now, https://pkg.go.dev seems not generate the library documentation automatically if the code contains generics.
I have written a small [gendoc](./gendoc) script and the generated document is available in [./doc.md](./doc.md).
Note that I don't really want to maintain this way and I hope generics official support in go package documentation.

## What to be implemented

- [x] ArrayList
- [ ] LinkedList
- [ ] Deque
- [ ] Queue
- [ ] PriorityQueue
- [ ] HashSet
- [ ] TreeSet
- [ ] Stack

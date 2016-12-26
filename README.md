# 👻 requirement
[![GoDoc](https://godoc.org/github.com/libkermit/requirement?status.png)](https://godoc.org/github.com/libkermit/requirement)
[![Build Status](https://travis-ci.org/libkermit/requirement.svg?branch=master)](https://travis-ci.org/libkermit/requirement)
[![Go Report Card](https://goreportcard.com/badge/github.com/libkermit/requirement)](https://goreportcard.com/report/github.com/libkermit/requirement)
[![License](https://img.shields.io/github/license/libkermit/requirement.svg)]()
[![codecov](https://codecov.io/gh/libkermit/requirement/branch/master/graph/badge.svg)](https://codecov.io/gh/libkermit/requirement)

`requirement` provides support for requirement to be able to skip some tests depending on the environment.
It is intended to work at lesat with the built-in `testing` framework. But any testing framework that
defines a `Skip(args ...interface{})` method can be used.

To write a test with a requirement, you need import this library and call it like the following

```go
    func TestWithRequirement(t *testing.T) {
        requirement.Is(t, func() bool {
            return false
        })
    }
```

This will print `requirement.go:51: requirement_test.go:56: unmatched requirement TestWithRequirement.func1`

## Provided requirement

A few requirement are also provided by this library. They are all based on the built-in libraries.

- `runtime` with `ArchitectureIs`, `OperatingSystemIs` and `GoVersionIs`.
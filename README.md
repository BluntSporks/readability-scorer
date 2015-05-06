# readability-scorer
Golang program to score readability using the readability package

## Purpose
To provide a program interface to the [readability](https://www.github.com/BluntSporks/readability) library which does
the actual calcuation of readability scores.

## Status
Ready to use.

However, this package has not been extensively tested. Results may differ in small amounts from official scoring of
these measures and should not be taken as official.

## Installation
This program is written in Google Go language. Make sure that Go is installed and the GOPATH is set up as described in
[How to Write Go Code](https://golang.org/doc/code.html).

The install this program and its dependencies by running:
    go install github.com/BluntSporks/readability-scorer

## Usage
readability-scorer [filename]

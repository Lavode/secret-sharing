# Introduction

This is a textbook implementation of t-out-of-n Shamir secret sharing, using
polynomials in a finite field of prime order.

This library was made for an assignment in class, and is not to be used
productively. Its interface only allows sharing integers, and it might well
be insecure.

# Getting started

Take a look at `demo.go` to see the library in use. If you've got a running
Golang setup, you may build & run it as follows:
```
go build demo.go && ./demo
```

# Project structure

The project structure is as follows:

* The `demo.go` application shows the library in use
* The `gf` package implements operations and polynomials over a finite field
* The `secretshare` package implements t-out-of-n secret sharing using
  polynomials of degree `t-1`

# Unit tests

Unit tests use the standard `testing` library of Go, and may be run using the
`go test` tool. To run all tests, execute `go test ./...`.

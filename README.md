[![Go Reference](https://pkg.go.dev/badge/github.com/kpurdon/zlg.svg)](https://pkg.go.dev/github.com/kpurdon/zlg)

*WORK IN PROGRESS*

zlg
---

zlg is an opinionated structured logger backed by https://github.com/rs/zerolog.

1. Enforces use of key:value fields over string formatted messages by excluding any formatting methods (e.g. `Errorf`, `Info`)
2. Ensures stack traces are always includeded with errors.
3. Provides quality default configuration of zerolog.
4. Abstracts typed field inclusion.

## Installation

```shell
go get github.com/kpurdon/zlg
```

## Quickstart

```go

// initialize a new logger
// for a default drop in use github.com/kpurdon/zlg/log
log := zlg.New()

// add fields to the logger instance
log.With("foo", "bar")

// write an error
log.Error(errors.New("some error"))

// add fields inline, and write a message at the info level
log.With("baz" 2).Info("hello")

// attach an error and panic
log.WithError(errors.New("some error")).Panic("hello")

```

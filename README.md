# go-notifications

[![Go](https://github.com/make-42/go-notifications/workflows/go/badge.svg)](https://github.com/make-42/go-notifications/actions?query=workflow%3Ago)
[![GoDoc](https://godoc.org/github.com/make-42/go-notifications?status.png)](https://godoc.org/github.com/make-42/go-notifications)
[![Go Report Card](https://goreportcard.com/badge/github.com/make-42/go-notifications)](https://goreportcard.com/badge/github.com/make-42/go-notifications)
[![codecov](https://codecov.io/gh/make-42/go-notifications/branch/main/graph/badge.svg)](https://codecov.io/gh/make-42/go-notifications)

go-notifications is an implementation of the notifications dbus interface written in go (golang) for receiving notifications.
Implemented and tested against version 1.2. See: https://specifications.freedesktop.org/notification-spec/1.2/.

## Example

An example using this library has been implemented and is accesible at `examples/example.go`.

```shell
git clone git@github.com:make-42/go-notifications.git

go build examples/example.go

./example
```

You can then test sending notifications with:

```shell
notify-send "Notification summary" "Notification body: test test test"
```

## Development

### Versioning

This library follows the semantic versioning concept.

### Commits

Commits should follow the conventional commit rules.  
See: https://conventionalcommits.org.

### Go Docs

Read the docs at https://pkg.go.dev/github.com/make-42/go-notifications

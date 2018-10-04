# go-starter

This is a set of simple exercises demonstrating basic Golang features explained in https://youtu.be/SqrbIlUwR0U and other notes for myself as I get more familiar with Go.

## Workspace

All code is organized into a **workspace**, where the location is defined in the `$GOPATH` env var. Its structure looks like this:

```
go
  bin
  src
    github.com
      jaredjensen
        go-starter
        other-project
    golang.org
  pkg
```

## Commands

| Command                 | Purpose                                             |
| ----------------------- | --------------------------------------------------- |
| `go version`            | Lists the version of Go installed                   |
| `go env`                | List all Go env vars                                |
| `go get {package}`      | Installs the specified package                      |
| `go run {path/file.go}` | Executes the Go file                                |
| `go build`              | Compiles the project to an executable               |
| `go install`            | Compiles and installs the executable to $GOPATH/bin |

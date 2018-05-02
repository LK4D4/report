## report

Tool to generate Go report that includes:
- Compiled code performance and binary sizes
- Go tool performance (compile time, etc.)
- Git logs for the period

## Usage

To generate report of Go changes included from `$rev1` to `$rev2`:

```
report -start=$rev1 -end=$rev2
```

By default, monitored packages are built in tmp folder under
isolated `GOPATH`. It's possible to set tested packages `GOPATH`
with `TEST_GOPATH` environment variable.

The directory reported by `go env GOROOT` (can be overriden by `GOROOT`
environment variable) is expected to be under git control, otherwise
it would not be possible for `report` to checkout specified revisions.

## Trivia

Generated reports are being discussed in [golangshow](http://golangshow.com/) podcast.
Create module:
```
go mod init example/hello/world
```

Installing module:
```
go install example/hello/world
go install .
go install 
```
Default folder is `~/go/bin/`
Discover install path:
```
go list -f '{{.Target}}'
```

Build:
`go build`

Adds missing requirements and remove requirements that are not longer used:
```
go mod tidy
```

Replace requirement with local copy
```
go mod edit -replace example/greetings=../greetings
```

Remove all downloaded modules:
```
go clean -modcache
```

Test package
```
go test
```

Test package with verbose output
```
go test -v
```


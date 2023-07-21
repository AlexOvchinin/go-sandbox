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

Build:
`go build`

Adds missing requirements and remove requirements that are not longer used:
```
go mod tidy
```

Remove all downloaded modules:
```
go clean -modcache
```

Test package
```
go test
```

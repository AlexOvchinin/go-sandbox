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

Add dependency:
```
go get golang.org/x/example
```

Test package
```
go test
```

Test package with verbose output
```
go test -v
```

Init workspace
```
go work init ./module-name
```

Additional commands for workspaces
```
go work use [-r] [dir] adds a use directive to the go.work file for dir, if it exists, and removes the use directory if the argument directory doesn’t exist. The -r flag examines subdirectories of dir recursively.
go work edit edits the go.work file similarly to go mod edit
go work sync syncs dependencies from the workspace’s build list into each of the workspace modules.
```
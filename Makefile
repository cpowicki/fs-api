SOURCES := $(shell find pkg -name '*.go' -not -path '*_test.go')

linux: dist/fs-api-linux-amd64
darwin: dist/fs-api-darwin-amd64
clean:
	rm dist/*

dist: 
	mkdir dist

dist/fs-api-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o dist/fs-api-linux-amd64 cmd/fs-api/main.go 

dist/fs-api-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o dist/fs-api-darwin-amd64 cmd/fs-api/main.go 
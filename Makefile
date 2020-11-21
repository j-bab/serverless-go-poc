.PHONY: build clean deploy

build:
	set GOOS=linux
	go build -ldflags="-s -w" -o bin/notes/create functions/notes/create.go
	go build -ldflags="-s -w" -o bin/notes/read functions/notes/read.go
	go build -ldflags="-s -w" -o bin/notes/update functions/notes/update.go
	go build -ldflags="-s -w" -o bin/notes/delete functions/notes/delete.go
	go build -ldflags="-s -w" -o bin/notes/list functions/notes/list.go

clean:
	rd /s /q  "./bin"

deploy: clean build
	sls deploy --verbose
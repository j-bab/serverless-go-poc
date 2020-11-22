.ONESHELL:

build:
	set GOOS=linux
	set GOARCH=amd64
	go build -ldflags="-s -w" -o bin/create functions/notes/create.go functions/notes/crudl.go
	go build -ldflags="-s -w" -o bin/read functions/notes/read.go functions/notes/crudl.go
	go build -ldflags="-s -w" -o bin/update functions/notes/update.go functions/notes/crudl.go
	go build -ldflags="-s -w" -o bin/delete functions/notes/delete.go functions/notes/crudl.go
	go build -ldflags="-s -w" -o bin/list functions/notes/list.go functions/notes/crudl.go

clean:
	if exist "./bin" rd /s /q  "./bin"

deploy: clean build
	sls deploy --verbose
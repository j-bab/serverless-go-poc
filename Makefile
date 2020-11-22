.ONESHELL:

build:
	set GOOS=linux
	set GOARCH=amd64
	go build -ldflags="-s -w" -o bin/notes functions/notes/main.go functions/notes/handlers.go functions/notes/crudl.go

clean:
	if exist "./bin" rd /s /q  "./bin"

deploy: clean build
	sls deploy --verbose
all: build

build:
	go build .

clean:
	rm -f liveplant-server

run: build
	PORT=9001 ./liveplant-server

heroku:
	godep go install && PORT=9001 foreman start

fmt:
	go fmt *.go

test: build
	go test .

coverage: build
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

install_deps:
	go get

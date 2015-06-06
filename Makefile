SOURCES = $(shell find . -name "*.go")

.PHONY: fmt print-% deps run clean

# TODO: test task and coverage

all: liveplant-server

liveplant-server: $(SOURCES)
	godep go build .

deps:
	godep save -r ./...

fmt:
	go fmt *.go

install: $(SOURCES)
	godep go install

run: liveplant-server
	foreman start

clean: 
	rm liveplant-server

print-%:
	@echo $*=$($*)

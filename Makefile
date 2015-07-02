SOURCES = $(shell find . -name "*.go")

.PHONY: all deps fmt run clean print-% 

# TODO: test task and coverage

all: liveplant-server

liveplant-server: $(SOURCES)
	godep go install ./...

deps:
	godep save -r ./...

fmt:
	go fmt *.go

run: liveplant-server
	foreman start

clean: 
	rm liveplant-server

print-%:
	@echo $*=$($*)

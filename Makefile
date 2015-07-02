SOURCES = $(shell find . -name "*.go")
TARGET = $(GOPATH)/bin/liveplant-server

.PHONY: all deps fmt run clean print-%

# TODO: test task and coverage

all: $(TARGET)

$(TARGET): $(SOURCES)
	godep go install ./...

deps:
	godep save -r ./...

fmt:
	go fmt *.go

run: liveplant-server
	foreman start

clean: 
	rm $(TARGET)

print-%:
	@echo $*=$($*)

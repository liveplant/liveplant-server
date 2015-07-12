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

run: $(TARGET)
	forego start

watch: $(TARGET)
	export LIVEPLANTDEBUG=1
	hr --conf hotreload.json

clean: 
	rm $(TARGET)

print-%:
	@echo $*=$($*)

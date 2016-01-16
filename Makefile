SOURCES_WEB = $(shell find ./liveplant-web/ -name "*.go")
TARGET_WEB = $(GOPATH)/bin/liveplant-web
SOURCES_CLOCK = $(shell find ./liveplant-clock/ -name "*.go")
TARGET_CLOCK = $(GOPATH)/bin/liveplant-clock
TARGET = $(TARGET_WEB) $(TARGET_CLOCK)

.PHONY: all deps fmt run clean print-%

# TODO: test task and coverage

all: $(TARGET)

$(TARGET_WEB): $(SOURCES_WEB)
	godep go install ./...

$(TARGET_CLOCK): $(SOURCES_CLOCK)
	godep go install ./...

deps:
	godep save -r ./...

fmt:
	go fmt *.go

run: $(TARGET)
	forego start

watch: $(TARGET)
	LIVEPLANTDEBUG=1 reflex -r '\.go$$' -s -- sh -c 'make run'

clean: 
	rm -f $(TARGET_WEB) $(TARGET_CLOCK)

print-%:
	@echo $*=$($*)

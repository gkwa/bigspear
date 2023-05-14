ifeq ($(OS),Windows_NT)
    SOURCES := $(shell dir /S /B *.go)
else
    SOURCES := $(shell find . -name '*.go')
endif

ifeq ($(shell uname),Darwin)
    GOOS = darwin
    GOARCH = amd64
    EXEEXT =
else ifeq ($(shell uname),Linux)
    GOOS = linux
    GOARCH = $(shell arch)
    EXEEXT =
else ifeq ($(shell uname),Windows_NT)
    GOOS = windows
    GOARCH = amd64
    EXEEXT = .exe
endif

TARGET := ./dist/bigspear_$(GOOS)_$(GOARCH)_v1/bigspear

run: bigspear
	./bigspear

$(TARGET): $(SOURCES)
	gofumpt -w $(SOURCES)
	goreleaser build --single-target --snapshot --clean
	go vet ./...


bigspear: $(TARGET)
	cp $< $@

all:
	goreleaser build --snapshot --clean

.PHONY: clean
clean:
	rm -f bigspear
	rm -f $(TARGET)
	rm -rf dist

BINARY_NAME := libcsvprocessor
GO_FILES := $(wildcard *.go)

# Regras
.PHONY: all build run clean test buildSO local

all: buildSO

build: $(GO_FILES)
	@go build -o $(BINARY_NAME) ./main.go

run: build
	@./$(BINARY_NAME)

clean:
	@go clean
	@rm -f $(BINARY_NAME)
	@rm -f libcsvprocessor.h libcsvprocess.so

test:
	@go test ./...

buildSO:
	@GCO_ENABLED=1 GOOS=linux \
		go build -buildmode=c-shared -o libcsvprocessor.so  \
		-ldflags '-s -w' \
		main.go
	@gcc -shared -o libcsv.so -fPIC libcsv.c -L. -lcsvprocessor -Wl,-rpath,./

local: 
	@gcc -o main main.c -L. -lcsv -Wl,-rpath,./ && ./main

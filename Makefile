.PHONY: all build install test clean example

PACKAGE_NAME=github.com/iures/daivplug
EXAMPLE_DIR=examples/basic-plugin
EXAMPLE_PLUGIN=basic-plugin.so

all: build

build:
	@echo "Building library..."
	@go build ./...

install:
	@echo "Installing library..."
	@go install ./...

test:
	@echo "Running tests..."
	@go test ./... -v

example:
	@echo "Building example plugin..."
	@cd $(EXAMPLE_DIR) && go build -buildmode=plugin -o $(EXAMPLE_PLUGIN)
	@echo "Example plugin built at $(EXAMPLE_DIR)/$(EXAMPLE_PLUGIN)"

clean:
	@echo "Cleaning up..."
	@rm -f $(EXAMPLE_DIR)/$(EXAMPLE_PLUGIN)
	@go clean

help:
	@echo "Available commands:"
	@echo "  make build    - Build the library"
	@echo "  make install  - Install the library"
	@echo "  make test     - Run tests"
	@echo "  make example  - Build example plugin"
	@echo "  make clean    - Clean up build artifacts" 

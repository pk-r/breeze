BINARY_NAME ?= breezed
SRC_DIR = ./cmd/$(BINARY_NAME)
PKG_DIR = ./...

# Go commands
GO_BUILD=go build
GO_CLEAN=go clean
GO_TEST=go test
GO_LDFLAGS=-X cmd/breezed.Version=0.0.1
GO_LDFLAGS=-ldflags "-X main.Version=0.0.1"

# Default target
all: build

# Build the binary
build:	 
	$(GO_BUILD) $(GO_LDFLAGS) -o $(BINARY_NAME) $(SRC_DIR)

# Run tests
test:
	$(GO_TEST) $(PKG_DIR)

clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

# Run the binary
run: build
	./$(BINARY_NAME)


help:
	@echo "Usage:"
	@echo "  make          - Build the binary"
	@echo "  make test     - Run tests"	
	@echo "  make clean    - Clean up generated files"
	@echo "  make run      - Build and run the binary"
	

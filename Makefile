# Define the binary name and the build directory
BINARY_NAME := fact
BUILD_DIR := bin

# Build rule: Compile the Go code and produce the executable
build:
	@go build -o $(BUILD_DIR)/$(BINARY_NAME)

# Run rule: Build the project first, then run the executable
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Test rule: Run tests with verbose output
test:
	go test -v ./...

# Clean rule: Remove the build artifacts
clean:
	rm -rf $(BUILD_DIR)/$(BINARY_NAME)

# PHONY targets to avoid conflicts with files named 'build', 'run', etc.
.PHONY: build run test clean
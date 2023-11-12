# Project-specific settings
BINARY_NAME := $(shell basename "$$PWD")
MAIN_GO := ./cmd/main.go # Define the path to your main Go file here

# Define phony targets to avoid conflicts with files of the same name and to improve performance
.PHONY: init install-dogo build watch

# Initial setup: install dogo, create config, and build the project
init: dogo-init dogo-config build gen-gitignore

# Install the dogo compiler for automatic rebuilds
dogo-init:
	go get github.com/liudng/dogo
	go install github.com/liudng/dogo
	go mod tidy

# Create a dogo.json configuration file if it doesn't exist
dogo-config:
	@if [ ! -e dogo.json ]; then \
		echo 'Creating default dogo.json configuration file...'; \
		echo '{' > dogo.json; \
		echo '    "WorkingDir": ".",' >> dogo.json; \
		echo '    "SourceDir": ["cmd"],' >> dogo.json; \
		echo '    "SourceExt": [".c", ".cpp", ".go", ".h"],' >> dogo.json; \
		echo '    "BuildCmd": "go build -o bin/$(BINARY_NAME) $(MAIN_GO)",' >> dogo.json; \
		echo '    "RunCmd": "./bin/$(BINARY_NAME)",' >> dogo.json; \
		echo '    "Decreasing": 1' >> dogo.json; \
		echo '}' >> dogo.json; \
	fi

# Clean the project: remove binary and clean Go cache
clean:
	@echo "  >  Cleaning build cache...\n"
	go clean
	rm -f $(BINARY_NAME)

# Build the application: compile the Go code in $(MAIN_GO) into a binary
build:
	@echo "  >  Building binary file...\n"
	go build -o bin/$(BINARY_NAME) $(MAIN_GO)

# Run the application: use dogo for automatic rebuilds on file changes
run:
	@echo "  >  Running application...\n"
	dogo -c dogo.json

# Generate a .gitignore file
gen-gitignore:
	@echo "  >  Generating .gitignore...\n"
	@echo "# Binaries" > .gitignore
	@echo "bin/" >> .gitignore
	@echo "# OS-specific files" >> .gitignore
	@echo "*.exe" >> .gitignore
	@echo "*.exe~" >> .gitignore
	@echo "*.dll" >> .gitignore
	@echo "*.so" >> .gitignore
	@echo "*.dylib" >> .gitignore
	@echo "# Test binary, built with 'go test -c'" >> .gitignore
	@echo "*.test" >> .gitignore
	@echo "# Output of the go coverage tool, specifically when used with LiteIDE" >> .gitignore
	@echo "*.out" >> .gitignore
	@echo "# Dependency directories (remove the comment below to include it)" >> .gitignore
	@echo "# vendor/" >> .gitignore

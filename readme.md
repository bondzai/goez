# Go-EZ-Toolbox

Go-EZ-Toolbox is a collection of Go utilities designed to simplify common programming tasks. This toolbox is intended for Go developers who want to speed up their development process with easy-to-use, reusable components.

## Features
- `PPrint`: Pretty-print JSON data in a human-readable format. Can filter to show only specified keys.
- `TrackPerformance`: Measure and log the execution time of code blocks, helpful for performance analysis.
- `Makefile (Optional)`: Setup & run the application, using dogo for automatic rebuilds on file changes.

## Installation
To use Go-EZ-Toolbox in your project, install it using `go get`:

```bash
go get github.com/introbond/go-ez-toolbox/toolbox
```

## Usage

### Pretty Printing JSON
```go
import "github.com/introbond/go-ez-toolbox/toolbox"

func main() {
    data := map[string]interface{}{
        "name": "Go-EZ-Toolbox",
        "purpose": "Simplify development",
    }

    // Pretty print all data
    toolbox.PPrint(data)

    // Pretty print specific keys
    toolbox.PPrint(data, "name")
}
```

### Tracking Performance
```go
import "github.com/introbond/go-ez-toolbox/toolbox"

func main() {
    defer toolbox.TrackPerformance("Example Task")()

    // Your code here
}
```

### (Optional) Cloning the Makefile

- Go to your Go project directory and install the Makefile.

```bash
curl -sL https://raw.githubusercontent.com/introbond/go-ez-toolbox/master/Makefile -o Makefile
```

- Makefile now valid on your project, and the content should look like this.

```bash
# Makefile for Golang dev.

# Project-specific settings
BINARY_NAME := $(shell basename "$$PWD")
MAIN_GO := ./cmd/main.go # Define the path to your main Go file here

# other code...
```

- After all run make init to setup boiler template & other command as you wish.

```bash
make init

# example for other cmds
make clean # run clean command in make file

make run # run the code with automatic rebuilds on file changes.
# etc...
```
- See the [Makefile](Makefile) file for more details.

# Go-EZ-Toolbox

Go-EZ-Toolbox is a collection of Go utilities designed to simplify common programming tasks. This toolbox is intended for Go developers who want to speed up their development process with easy-to-use, reusable components.

## Features
- `Makefile`: Setup & run the application, using dogo for automatic rebuilds on file changes.
- `PPrint`: Pretty-print JSON data in a human-readable format. Can filter to show only specified keys.
- `TrackPerformance`: Measure and log the execution time of code blocks, helpful for performance analysis.

## Installation
To use Go-EZ-Toolbox in your project, install it using `go get`:

```bash
go get github.com/introbond/go-ez-toolbox/toolbox
```

## Usage

### Cloning the Makefile
```bash
curl -sL https://raw.githubusercontent.com/introbond/go-ez-toolbox/master/Makefile -o Makefile
```

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

# Chronus

[![Go Version](https://img.shields.io/github/go-mod/go-version/reinanbr/chronus)](https://golang.org/)
[![License](https://img.shields.io/github/license/reinanbr/chronus)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/reinanbr/chronus)](https://goreportcard.com/report/github.com/reinanbr/chronus)
[![GoDoc](https://godoc.org/github.com/reinanbr/chronus?status.svg)](https://godoc.org/github.com/reinanbr/chronus)

A simple and lightweight Go library for time utilities, providing millisecond-precision time operations.

## Features

- Get current time in milliseconds since Unix epoch
- Sleep functionality with millisecond precision
- Simple and intuitive API
- Zero dependencies beyond Go standard library

## Installation

```bash
go get github.com/reinanbr/chronus
```

## Usage

### Getting Current Time

```go
package main

import (
    "fmt"
    "github.com/reinanbr/chronus"
)

func main() {
    // Get current time in milliseconds since Unix epoch
    now := chronus.Now()
    fmt.Printf("Current time: %d milliseconds\n", now)
}
```

### Sleep Functionality

```go
package main

import (
    "fmt"
    "time"
    "github.com/reinanbr/chronus"
)

func main() {
    fmt.Println("Starting...")
    
    // Sleep for 1000 milliseconds (1 second)
    chronus.Sleep(1000)
    
    fmt.Println("Finished after 1 second")
}
```

### Complete Example

```go
package main

import (
    "fmt"
    "github.com/reinanbr/chronus"
)

func main() {
    start := chronus.Now()
    fmt.Printf("Start time: %d\n", start)
    
    // Sleep for 2 seconds
    chronus.Sleep(2000)
    
    end := chronus.Now()
    fmt.Printf("End time: %d\n", end)
    
    elapsed := end - start
    fmt.Printf("Elapsed time: %d milliseconds\n", elapsed)
}
```

## API Reference

### `Now() int64`

Returns the current time in milliseconds since the Unix epoch (January 1, 1970 UTC).

**Returns:**
- `int64`: Current time in milliseconds

### `Sleep(duration time.Duration)`

Pauses execution for the specified duration in milliseconds.

**Parameters:**
- `duration time.Duration`: Duration to sleep in milliseconds

## Examples

You can find a complete example in the `example/` directory. To run it:

```bash
cd example
go run main.go
```

## Development

### Running Tests

```bash
make test
```

### Running Examples

```bash
make example
```

### Building

```bash
make build
```

### Linting

```bash
make lint
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

### v1.0.0
- Initial release
- Added `Now()` function for getting current time in milliseconds
- Added `Sleep()` function for millisecond-precision sleeping

## Documentation

- [🏗️ Architecture & Structure](ARQUITETURA_ESTRUTURA.md) - Detailed explanation of library structure and why each part is necessary
- [📚 Publication Guide](GUIA_PUBLICACAO.md) - Complete guide on how to publish a Go library to pkg.go.dev
- [🎯 Chronus Case Study](CASO_PRATICO_CHRONUS.md) - Practical example of this library's publication process

## Support

If you have any questions or need help, please open an issue on GitHub.

# genv

## Ergonomic environment management

<code>genv</code> is a Go package for loading and accessing environment variables. It simplifies the processes of lookup, casting, and populating configuration structs, using generics, struct tags, and reflection for mapping.

## Features

- Load and cast environment variables through the use of generics.
- Supports basic types like <code>string</code>, <code>bool</code>, <code>int</code>, <code>float64</code>, etc.
- Load directly into a struct, including nested structs for more complex configurations.
- Autoload (.env only) via import _ "github.com/brendanjcarlson/genv/autoload"

## Installation

```bash
go get -u github.com/brendanjcarlson/genv
```

## Usage

```go
package main

import "github.com/brendanjcarlson/genv"

func main() {
    err := genv.Load(".env")
    if err != nil {
        log.Fatalf("load env: %v", err)
    }

    superSecretKey, err := genv.Get[string]("SUPER_SECRET_KEY")
    if err != nil {
        log.Fatalln(err)
    }

    timeoutSeconds, err := genv.Get[int]("TIMEOUT_SECONDS")
    if errors.Is(err, genv.ErrCannotCast) {
        log.Println(err)
    } else {
        timeoutSeconds = 5
    }

    mustHaveValue := genv.GetOrPanic[string]("GOTTA_HAVE_THIS_ONE")

    type ServerConfig struct {
        Host string `genv:"SERVER_HOST"`
        Port string `genv:"SERVER_PORT"`
    }

    var serverConfig ServerConfig
    if err := genv.GetStruct(&serverConfig); err !=nil {
        log.Fatalln(err)
    }
}
```

## Supported Types

The following types are currently supported, with support for slices coming:

- <code>string</code>
- <code>bool</code>
- <code>int</code>, <code>int8</code>, <code>int16</code>, <code>int32</code>, <code>int64</code>
- <code>uint</code>, <code>uint8</code>, <code>uint16</code>, <code>uint32</code>, <code>uint64</code>
- <code>float32</code>, <code>float64</code>
- Structs with the listed types (including nested structs)

## Documentation

See user documentation at <https://pkg.go.dev/github.com/brendanjcarlson/genv>

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request on GitHub.

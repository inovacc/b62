# B62 is a base62 encoding and decoding messages

B62 is a Go package for encoding and decoding data using Base62 encoding. Base62 encoding is a method of representing binary data in an ASCII string format using 62 characters (A-Z, a-z, 0-9). This package provides functions to encode and decode data efficiently.

## Features

- Encode binary data to Base62 string
- Decode Base62 string back to binary data
- Handles edge cases and invalid characters gracefully

## Installation

To install the package, use the following command:

```sh
go get github.com/inovacc/b62
```

# Encode binary data to Base62 string

```go
package main

import (
    "fmt"
    "github.com/inovacc/b62"
)

func main() {
    data := []byte("Hello, World!")
    encoded := b62.Encode(data)
    fmt.Println("Encoded:", encoded)
}
```

# Decode Base62 string back to binary data

```go
package main

import (
    "fmt"
    "github.com/inovacc/b62"
)

func main() {
    encoded := "Base62EncodedString"
    decoded, err := b62.Decode(encoded)
    if err != nil {
        fmt.Println("Error decoding:", err)
        return
    }
    fmt.Println("Decoded:", string(decoded))
}
```
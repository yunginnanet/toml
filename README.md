# toml

`github.com/yunginnanet/toml`

Go module for marshaling and unmarshaling TOML data with strict schema enforcement via struct tags.

## Features

- Encode Go structs into TOML
- Decode TOML into nested Go structs
- Schema validation using `toml` struct tags
- Supports nested tables, pointer sub-structs, slices
- Ignores unexported or explicitly skipped fields

## Usage

### Encoding

Structs must be composed entirely of exported sub-structs, each with valid `toml` tags.

```go
type Config struct {
	Server ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Port int    `toml:"port"`
	Host string `toml:"host"`
}

data, err := toml.MarshalTOML(Config{
	Server: ServerConfig{Port: 8080, Host: "localhost"},
})
```

Produces:

```go
[server]
port = 8080
host = "localhost"
```

### Decoding

The same struct layout is used for unmarshaling.

Example:

```go
data := `[server]
port = 8080
host = "localhost"`

input := []byte(data)

var conf Config
err := toml.UnmarshalTOML(input, &conf)
```

### Notes

- Fields without a `toml` tag are ignored.
- Fields tagged with `toml:"-"` are explicitly skipped.
- Empty sub-structs are omitted from output.
- Only structs and nested structs are supported at the top level.

## Errors

Returns descriptive errors for:

- Missing or malformed TOML syntax
- Structs missing required `toml` tags
- Non-pointer or non-struct targets for decoding
- Type mismatches between TOML values and Go fields

## Tests

Extensive tests cover:

- Basic and nested struct serialization
- Pointer and value struct decoding
- Field skipping
- Invalid schema handling

Run tests with:

    go test ./...

---

Built with pool-backed buffer management from [github.com/yunginnanet/common/pool](https://github.com/yunginnanet/common/pool).

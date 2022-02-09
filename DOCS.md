# Documentation

To use this package, first install it in your project by running:

```bash
go get github.com/hisamafahri/lodang
```

## Usage

### Slices

- Chunk

```go
slice := []interface{}{1, 2, 3, 4}
result := lodang.Chunk(slice, 3)

// result: [[1 2 3] [4]]
// type: []interface{}
```

## License

[MIT](LICENSE)
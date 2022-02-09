# Documentation

### Slices

- Chunk

Creates a `slice` of elements split into groups the length of `size`. If the `slice` can't be split evenly, the final chunk will be the remaining elements.

```go
slice := []interface{}{1, 2, 3, 4}
result := lodang.Chunk(slice, 3)

// result: [[1 2 3] [4]]
// type: []interface{}
```

## License

[MIT](LICENSE)
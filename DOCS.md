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

- Compact

Creates a `slice` with all falsey values removed. The values `false`, `null`, `0`, `""`, `undefined`, and `NaN` are falsey

```go
slice := []interface{}{0, false, nil, "", "a", 1, 2, 3, 4}
result := lodang.Compact(slice)

// result: [1 2 3 4]
// type: []interface{}
```

- Compact

Creates a new `slice` concatenating `slice` with any additional `slices`. Currently only support **two** `[]interface{}` type format.

```go
a := []interface{}{9}
b := []interface{}{1, 2}
result := lodang.Concat(a, b)

// result: [9 1 2]
// type: []interface{}
```

## License

[MIT](LICENSE)
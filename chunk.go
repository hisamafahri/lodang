package lodang

/**
 * Creates a slice of elements split into groups the length of `size`.
 * If `slize` can't be split evenly, the final chunk will be the remaining
 * elements.
 *
 * @example
 *
 * Chunk([]interface{}{'a', 'b', 'c', 'd'}, 2)
 * // => [[]interface{}{'a', 'b'}, []interface{}{'c', 'd'}]
 *
 * chunk([]interface{}{'a', 'b', 'c', 'd'{}, 3)
 * // => [[]interface{}{'a', 'b', 'c'}, []interface{}{'d'}]
 */

func Chunk(slice []interface{}, size uint) []interface{} {
	if len(slice) == 0 {
		return nil
	}

	divided := make([]interface{}, (len(slice)+int(size)-1)/int(size))
	prev := 0
	i := 0
	till := len(slice) - int(size)

	for prev < till {
		next := prev + int(size)
		divided[i] = slice[prev:next]
		prev = next
		i++
	}

	divided[i] = slice[prev:]
	return divided
}

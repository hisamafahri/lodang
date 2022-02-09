package lodang

/**
 * Creates a new `slice` concatenating `slice` with any additional `slices`.
 * Currently only support **two** `[]interface{}` type format.
 * @example
 *
 * concat([]interface{}{9}, []interface{}{1, 2})
 * // => [9 1 2]
 */

func Concat(slice []interface{}, newSlices []interface{}) []interface{} {
	return append(slice, newSlices...)
}

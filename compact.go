package lodang

import (
	"reflect"
)

/**
 * Creates a `slice` with all falsey values removed. The values `false`, `null`,
 * `0`, `""`, `undefined`, and `NaN` are falsey.
 * @example
 *
 * compact([]interface{}{0, 1, false, 2, '', 3})
 * // => [1, 2, 3]
 */

func Compact(slice []interface{}) []interface{} {

	var dashSlice []interface{}

	for _, item := range slice {
		if !reflect.DeepEqual(item, nil) &&
			!reflect.DeepEqual(item, false) &&
			!reflect.DeepEqual(item, 0) &&
			!reflect.DeepEqual(item, "") {
			dashSlice = append(dashSlice, item)
		}
	}

	return dashSlice
}

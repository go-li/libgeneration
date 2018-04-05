// Package binary implements binary heap
package binary

func percolateUp(array *[], hole int, val *, cmp func(*, *) int) (next int) {
	for {
		next = ((hole + 1) >> 1) - 1
		if !((hole > 0) && (cmp(val, &(*array)[next]) < 0)) {
			break
		}
		var l = &(*array)[hole]
		var r = &(*array)[next]
		*l = *r
		hole = next
	}
	return hole
}

func Insert(array *[], add [], cmp func(*, *) int) {
	if len(*array)+len(add) <= cap(*array) {
		*array = (*array)[0 : len(*array)+len(add)]
	} else if cap(*array) > 0 {
		*array = append(*array, add...)
	} else {
		*array = make([], len(add))
	}
	for i := range add {
		var q = &add[i]
		var p = &(*array)[percolateUp(array, len(*array)-1, q, cmp)]

		*p = *q
	}
}

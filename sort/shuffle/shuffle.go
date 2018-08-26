package shuffle

import "math/rand"

func Order(slice [], random func(int)int) {
	var temp *
	if len(slice) >= cap(slice) {
		var tmp []
		tmp = make([], 1)
		temp = &tmp[0]
	} else {
		temp = &(slice[0 : len(slice)+1])[len(slice)]
	}
	for len(slice) > 0 {
		n := len(slice)
		randIndex := random(n)
		var last *;
		var any *;
		last = &slice[n-1]
		any = &slice[randIndex]

		*temp = *last
		*last = *any
		*any = *temp
		slice = slice[:n-1]
	}
}

func Random(slice []) {
	Order(slice, func(n int) int {
		return rand.Intn(n)
	})
}

func Rand(slice [], rand *rand.Rand) {
	Order(slice, func(n int) int {
		return rand.Intn(n)
	})
}

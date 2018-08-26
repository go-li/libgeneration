package set

func Find(get func(*)*, set func(*,*), cmp func(*,*) int, i *) (j *) {
	j = i
	for cmp(j, get(j)) != 0 {
		j = get(j)
	}
	for cmp(i, get(i)) != 0 {
		var tmp = get(i)
		set(i,j)
		i = tmp
	}
	return j
}

func Union(get func(*)*, set func(*,*), cmp func(*,*) int, a *, b *) {

	a = Find(get, set, cmp, a)
	b = Find(get, set, cmp, b)

	var c = cmp(a, b)
	if c < 0 {
		set(a, a)
		set(b, a)
	}
	if c > 0 {
		set(b, b)
		set(a, b)
	}
}

package graph

func adaptive_sort(a [], tmp *, cmp func(*, *) int) {
	for i := 1; i < len(a); i++ {
		var p = &a[i]
		*tmp = *p
		j := i - 1
		for ; j >= 0 && cmp(&a[j], tmp) > 0; j-- {
			var q = &a[j+1]
			var r = &a[j]
			*q = *r
		}
		p = &a[j+1]
		*p = *tmp
	}
}

func sorted_sorted_delete(s [], deletion [], cmp func(*,*) int) [] {
	for i := 0; i < len(s); i++ {
		for (len(deletion) > 0) && cmp(&s[i], &deletion[0]) == 0 {
			enddeleted := true
			for len(deletion) > 0 && enddeleted {
				enddeleted = false
				ed := &deletion[len(deletion)-1]

				for (len(s) > 0) && cmp(&s[len(s)-1], ed) == 0 {
					s = s[0:len(s)-1]
					enddeleted = true
				}
				deletion = deletion[0:len(deletion)-1]

			}

			if len(s) <= i {
				break
			}

			l := &s[i]
			r := &s[len(s)-1]

			*l = *r
			s = s[0:len(s)-1]
		}

		if len(s) <= i {
			break
		}

		for (len(deletion) > 0) && (cmp(&s[i], &deletion[0]) > 0) {
			deletion = deletion[1:]
		}
	}
	return s
}

func cmp_ints(a, b *int) int {
	return *a - *b
}

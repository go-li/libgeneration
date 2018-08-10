package avl

func linker(links *[3]*, _ *byte) *[3]* {
	return links
}

func balancer(_ *[3]*, balance *byte) *byte {
	return balance
}

func Append(tree func(*) (*[3]*, *byte), root **, key *, cmp func(*,*) int) * {
	var y * ;
	var p * ;
	var q * ;
	var n * ;
	var w * ;

	var dir int

	y = *root

	q = nil
	for p = *root; p != nil; {
		var compare int = cmp(key, p)
		if compare == 0 {
			return p
		}
		if compare > 0 {
			dir = 1
		} else {
			dir = 0
		}

		if *balancer(tree(p)) != 0 {
			y = p
		}

		q = p
		p = linker(tree(p))[dir]
	}

	n = key

	linker(tree(n))[0] = nil
	linker(tree(n))[1] = nil
	linker(tree(n))[2] = q

	if q != nil {
		linker(tree(q))[dir] = n
	} else {
		*root = n
	}
	*balancer(tree(n)) = 0
	if *root == n {
		return n
	}

	for p = n; p != y; p = q {
		q = linker(tree(p))[2]
		if linker(tree(q))[0] != p {
			dir = 1
		} else {
			dir = 0
		}
		if dir == 0 {
			*balancer(tree(q))--
		} else {
			*balancer(tree(q))++
		}
	}

	if *balancer(tree(y)) == 254 {
		var x *;
		x = linker(tree(y))[0]
		if *balancer(tree(x)) == 255 {
			w = x
			linker(tree(y))[0] = linker(tree(x))[1]
			linker(tree(x))[1] = y
			*balancer(tree(x)) = 0
			*balancer(tree(y)) = 0
			linker(tree(x))[2] = linker(tree(y))[2]
			linker(tree(y))[2] = x
			if linker(tree(y))[0] != nil {
				linker(tree(linker(tree(y))[0]))[2] = y
			}
		} else {
			if !(*balancer(tree(x)) == +1) {
				panic("assert")
			}
			w = linker(tree(x))[1]
			linker(tree(x))[1] = linker(tree(w))[0]
			linker(tree(w))[0] = x
			linker(tree(y))[0] = linker(tree(w))[1]
			linker(tree(w))[1] = y
			if *balancer(tree(w)) == 255 {
				*balancer(tree(x)) = 0
				*balancer(tree(y)) = +1
			} else if *balancer(tree(w)) == 0 {
				*balancer(tree(x)) = 0
				*balancer(tree(y)) = 0
			} else { /* |balancer(tree(w))[31] == +1| */
				*balancer(tree(x)) = 255
				*balancer(tree(y)) = 0
			}
			*balancer(tree(w)) = 0
			linker(tree(w))[2] = linker(tree(y))[2]
			linker(tree(x))[2] = w
			linker(tree(y))[2] = w
			if linker(tree(x))[1] != nil {
				linker(tree(linker(tree(x))[1]))[2] = x
			}
			if linker(tree(y))[0] != nil {
				linker(tree(linker(tree(y))[0]))[2] = y
			}
		}
	} else if *balancer(tree(y)) == +2 {
		var x *;
		x = linker(tree(y))[1]
		if *balancer(tree(x)) == +1 {
			w = x
			linker(tree(y))[1] = linker(tree(x))[0]
			linker(tree(x))[0] = y
			*balancer(tree(x)) = 0
			*balancer(tree(y)) = 0
			linker(tree(x))[2] = linker(tree(y))[2]
			linker(tree(y))[2] = x
			if linker(tree(y))[1] != nil {
				linker(tree(linker(tree(y))[1]))[2] = y
			}
		} else {
			if !(*balancer(tree(x)) == 255) {
				panic("assert")
			}
			w = linker(tree(x))[0]
			linker(tree(x))[0] = linker(tree(w))[1]
			linker(tree(w))[1] = x
			linker(tree(y))[1] = linker(tree(w))[0]
			linker(tree(w))[0] = y
			if *balancer(tree(w)) == +1 {
				*balancer(tree(x)) = 0
				*balancer(tree(y)) = 255
			} else if *balancer(tree(w)) == 0 {
				*balancer(tree(x)) = 0
				*balancer(tree(y)) = 0
			} else { /* |balancer(tree(w)) == 255| */
				*balancer(tree(x)) = +1
				*balancer(tree(y)) = 0
			}
			*balancer(tree(w)) = 0
			linker(tree(w))[2] = linker(tree(y))[2]
			linker(tree(x))[2] = w
			linker(tree(y))[2] = w
			if linker(tree(x))[0] != nil {
				linker(tree(linker(tree(x))[0]))[2] = x
			}
			if linker(tree(y))[1] != nil {
				linker(tree(linker(tree(y))[1]))[2] = y
			}
		}
	} else {
		return n
	}
	if linker(tree(w))[2] != nil {
		var oo = 0
		if y != linker(tree(linker(tree(w))[2]))[0] {
			oo = 1
		}
		linker(tree(linker(tree(w))[2]))[oo] = w
	} else {
		*root = w
	}

	return n
}


/* Deletes from |tree| and returns an item matching |item|.
   Returns a null pointer if no matching item found. */
func Delete(tree func(*) (*[3]*, *byte), root **, key *, cmp func(*,*) int) (result *) {

	var p *;
	var q *; 
	var dir int

	if !(tree != nil) {
		panic("assert")
	}

	if *root == nil {
		return nil
	}

	p = *root
	for {
		var compare int = cmp(key, p)
		if compare == 0 {
			break
		}

		if compare > 0 {
			dir = 1
		} else {
			dir = 0
		}

		p = linker(tree(p))[dir]
		if p == nil {
			return nil
		}
	}
	result = p

	q = linker(tree(p))[2]
	if q == nil {
		q = *root
		dir = 0
	}

	if linker(tree(p))[1] == nil {
		linker(tree(q))[dir] = linker(tree(p))[0]
		if linker(tree(q))[dir] != nil {
			linker(tree(linker(tree(q))[dir]))[2] = linker(tree(p))[2]
		}
	} else {
		var r *;
		r = linker(tree(p))[1]
		if linker(tree(r))[0] == nil {
			linker(tree(r))[0] = linker(tree(p))[0]
			linker(tree(q))[dir] = r
			linker(tree(r))[2] = linker(tree(p))[2]
			if linker(tree(r))[0] != nil {
				linker(tree(linker(tree(r))[0]))[2] = r
			}
			*balancer(tree(r)) = *balancer(tree(p))
			q = r
			dir = 1
		} else {
			var s *;
			s = linker(tree(r))[0]
			for linker(tree(s))[0] != nil {
				s = linker(tree(s))[0]
			}
			r = linker(tree(s))[2]
			linker(tree(r))[0] = linker(tree(s))[1]
			linker(tree(s))[0] = linker(tree(p))[0]
			linker(tree(s))[1] = linker(tree(p))[1]
			linker(tree(q))[dir] = s
			if linker(tree(s))[0] != nil {
				linker(tree(linker(tree(s))[0]))[2] = s
			}
			linker(tree(linker(tree(s))[1]))[2] = s
			linker(tree(s))[2] = linker(tree(p))[2]
			if linker(tree(r))[0] != nil {
				linker(tree(linker(tree(r))[0]))[2] = r
			}
			*balancer(tree(s)) = *balancer(tree(p))
			q = r
			dir = 0
		}
	}

	if p == *root && p != nil {
		if linker(tree(p))[0] == nil {
			if q == *root {
				q = linker(tree(p))[1]	
			}
			*root = linker(tree(p))[1]
		} else {
			if q == *root {
				q = linker(tree(p))[0]	
			}
			*root = linker(tree(p))[0]
		}
	}

	linker(tree(p))[0] = nil
	linker(tree(p))[1] = nil
	linker(tree(p))[2] = nil
	*balancer(tree(p)) = 0

	//  tree.pavl_alloc.libavl_free (tree.pavl_alloc, p);

	for q != *root {
		var y *;
		y = q

		if linker(tree(y))[2] != nil {
			q = linker(tree(y))[2]
		} else {
			q = *root
		}

		if dir == 0 {
			if linker(tree(q))[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			*balancer(tree(y))++
			if *balancer(tree(y)) == +1 {
				break
			} else if *balancer(tree(y)) == +2 {
				var x *;
				x = linker(tree(y))[1]
				if *balancer(tree(x)) == 255 {
					var w *;

					if !(*balancer(tree(x)) == 255) {
						panic("assert")
					}
					w = linker(tree(x))[0]
					linker(tree(x))[0] = linker(tree(w))[1]
					linker(tree(w))[1] = x
					linker(tree(y))[1] = linker(tree(w))[0]
					linker(tree(w))[0] = y
					if *balancer(tree(w)) == +1 {
						*balancer(tree(x)) = 0
						*balancer(tree(y)) = 255
					} else if *balancer(tree(w)) == 0 {
						*balancer(tree(x)) = 0
						*balancer(tree(y)) = 0
					} else { /* |balancer(tree(w)) == 255| */
						*balancer(tree(x)) = +1
						*balancer(tree(y)) = 0
					}
					*balancer(tree(w)) = 0
					linker(tree(w))[2] = linker(tree(y))[2]
					linker(tree(x))[2] = w
					linker(tree(y))[2] = w
					if linker(tree(x))[0] != nil {
						linker(tree(linker(tree(x))[0]))[2] = x
					}
					if linker(tree(y))[1] != nil {
						linker(tree(linker(tree(y))[1]))[2] = y
					}
					linker(tree(q))[dir] = w
				} else {
					linker(tree(y))[1] = linker(tree(x))[0]
					linker(tree(x))[0] = y
					linker(tree(x))[2] = linker(tree(y))[2]
					linker(tree(y))[2] = x
					if linker(tree(y))[1] != nil {
						linker(tree(linker(tree(y))[1]))[2] = y
					}
					linker(tree(q))[dir] = x
					if *balancer(tree(x)) == 0 {
						*balancer(tree(x)) = 255
						*balancer(tree(y)) = +1
						break
					} else {
						*balancer(tree(x)) = 0
						*balancer(tree(y)) = 0
						y = x
					}
				}
			}
		} else {
			if linker(tree(q))[0] != y {
				dir = 1
			} else {
				dir = 0
			}
			*balancer(tree(y))--
			if *balancer(tree(y)) == 255 {
				break
			} else if *balancer(tree(y)) == 254 {
				var x *;
				x = linker(tree(y))[0]
				if *balancer(tree(x)) == +1 {
					var w *;
					if !(*balancer(tree(x)) == +1) {
						panic("assert")
					}
					w = linker(tree(x))[1]
					linker(tree(x))[1] = linker(tree(w))[0]
					linker(tree(w))[0] = x
					linker(tree(y))[0] = linker(tree(w))[1]
					linker(tree(w))[1] = y
					if *balancer(tree(w)) == 255 {
						*balancer(tree(x)) = 0
						*balancer(tree(y)) = +1
					} else if *balancer(tree(w)) == 0 {
						*balancer(tree(x)) = 0
						*balancer(tree(y)) = 0
					} else { /* |balancer(tree(w)) == +1| */
						*balancer(tree(x)) = 255
						*balancer(tree(y)) = 0
					}
					*balancer(tree(w)) = 0
					linker(tree(w))[2] = linker(tree(y))[2]
					linker(tree(x))[2] = w
					linker(tree(y))[2] = w
					if linker(tree(x))[1] != nil {
						linker(tree(linker(tree(x))[1]))[2] = x
					}
					if linker(tree(y))[0] != nil {
						linker(tree(linker(tree(y))[0]))[2] = y
					}
					linker(tree(q))[dir] = w
				} else {
					linker(tree(y))[0] = linker(tree(x))[1]
					linker(tree(x))[1] = y
					linker(tree(x))[2] = linker(tree(y))[2]
					linker(tree(y))[2] = x
					if linker(tree(y))[0] != nil {
						linker(tree(linker(tree(y))[0]))[2] = y
					}
					linker(tree(q))[dir] = x
					if *balancer(tree(x)) == 0 {
						*balancer(tree(x)) = +1
						*balancer(tree(y)) = 255
						break
					} else {
						*balancer(tree(x)) = 0
						*balancer(tree(y)) = 0
						y = x
					}
				}
			}
		}
	}
	return result

}

func Foreach(tree func(*) (*[3]*, *byte), root *, apply func(*)bool) bool {

	var paths, _ = tree(root)

	if (paths[0] != nil) && !Foreach(tree, paths[0], apply) {
		return false
	}
	if !apply(root) {
		return false
	}
	if (paths[1] != nil) && !Foreach(tree, paths[1], apply) {
		return false
	}
	return true
}

func Find(tree func(*) (*[3]*, *byte), root *, find func(*) byte) (f *) {
	f = root
	for {
		var children, _ = tree(f)
		var dir byte
		if find != nil {
			dir = find(f)
			if dir > 3 {
				if dir >= 7 {
					return nil
				}
				if dir == 6 {
					return f
				}
				return (*children)[dir-3]
			}
		}

		if (*children)[dir] == nil {
			return f
		}
		f = (*children)[dir]
	}
}

func InOrderSuccessor(tree func(*) (*[3]*, *byte), n *) * {
	if( linker(tree(n))[1] != nil ) {
		return Find(tree, linker(tree(n))[1], nil);
	}

	var p * ;
	p = linker(tree(n))[2];
	for p != nil && n == linker(tree(p))[1]  {
		n = p;
		p = linker(tree(p))[2];
	}
	return p;
}

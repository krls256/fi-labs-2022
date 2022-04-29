package ring

type Ring struct {
	mod int
}

func NewRing(mod int) (*Ring, error) {
	if mod <= 1 {
		return nil, ModuleIsNotAdmissibleErr
	}
	return &Ring{mod: mod}, nil
}
func (r *Ring) Mod() int {
	return r.mod
}

func (r *Ring) Inverse(a int) (int, error) {
	a = a % r.mod
	r1, r2 := r.mod, a
	qs := []int{}
	for r2 != 0 {
		qs = append(qs, r1/r2)
		r1, r2 = r2, r1%r2
	}
	if r1 == 1 {
		s := []int{0, 1}
		for i, j := len(qs)-2, 1; i >= 0; i-- {
			s = append(s, s[j-1]-qs[i]*s[j])
			j++
		}
		return (s[len(s)-1] + r.mod) % r.mod, nil
	}
	return 0, InverseElementIsNotExistErr
}

func gcd(a, b int) int {
	if b > a {
		return gcd(b, a)
	}
	r1, r2 := a, b
	for r2 != 0 {
		r1, r2 = r2, r1%r2
	}
	return r1
}

// ax === b (mod n)
func (r *Ring) SolveCongruence(a, b int) []int {
	d := gcd(r.mod, a)
	if d == 1 {
		aN, err := r.Inverse(a)
		if err != nil {
			panic("inverse works incorrect")
		}
		return []int{(b * aN) % r.mod}
	} else if b%d == 0 {
		b1 := b / d
		n1 := r.mod / d
		a1 := a / d
		newR, err := NewRing(n1)
		if err != nil {
			panic("impossible")
		}
		a1N, err := newR.Inverse(a1)
		if err != nil {
			panic("inverse works incorrect")
		}
		x0 := b1 * a1N % newR.mod
		res := []int{}
		for i := 0; i < d; i++ {
			res = append(res, x0+i*n1)
		}
		return res
	}
	return []int{}
}

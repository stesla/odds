package odds

func Interpret(diff int, dice []int) (r int, botch bool) {
	var nb, ns int

	for _, d := range dice {
		if d == 1 {
			nb++
		} else if d >= diff {
			ns++
		}
	}

	if ns > 0 {
		if i := ns - nb; i > 0 {
			r = i
		}
	} else if nb > 0 {
		botch = true
	}

	return
}

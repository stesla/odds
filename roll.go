package main

import "math/rand"

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

func Roll(r *rand.Rand, n int) (dice []int) {
	for i := 0; i < n; i++ {
		dice = append(dice, r.Intn(10)+1)
	}
	return
}

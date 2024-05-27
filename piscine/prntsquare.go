package piscine

import "github.com/01-edu/z01"

func printSquare(c, r int, tl, tr, bl, br, m, h, v rune) {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 {
				z01.PrintRune(tl)
			} else if i == 0 && j == c-1 {
				z01.PrintRune(tr)
			} else if i == r-1 && j == 0 {
				z01.PrintRune(bl)
			} else if i == r-1 && j == c-1 {
				z01.PrintRune(br)
			} else if i == 0 || i == r-1 {
				z01.PrintRune(h)
			} else if j == 0 || j == c-1 {
				z01.PrintRune(v)
			} else {
				z01.PrintRune(m)
			}
		}
		z01.PrintRune('\n')
	}
}

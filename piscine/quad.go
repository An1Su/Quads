package piscine

func QuadA(x, y int) {
	printSquare(x, y, 'o', 'o', 'o', 'o', ' ', '-', '|')
}

func QuadB(x, y int) {
	printSquare(x, y, '/', '\\', '\\', '/', ' ', '*', '*')
}

func QuadC(x, y int) {
	printSquare(x, y, 'A', 'A', 'C', 'C', ' ', 'B', 'B')
}

func QuadD(x, y int) {
	printSquare(x, y, 'A', 'C', 'A', 'C', ' ', 'B', 'B')
}

func QuadE(x, y int) {
	printSquare(x, y, 'A', 'C', 'C', 'A', ' ', 'B', 'B')
}

package quad

import "fmt"

func QuadA(x int, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if (row == 0 || row == y-1) && (col == 0 || col == x-1) {
				fmt.Print("o")
			} else if row == 0 || row == y-1 {
				fmt.Print("-")
			} else if col == 0 || col == x-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

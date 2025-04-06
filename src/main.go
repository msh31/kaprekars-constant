package main

import (
	"fmt"
	"sort"
)

func isValidNumber(n int) bool {
	return n >= 1000 && n <= 9998
}

func main() {
	var n, count, result int

	fmt.Print("\033[H\033[2J") // Clear the console (thanks stackoverflow)

	for {
		fmt.Print("Enter a 4-digit number: ")
		var _, err = fmt.Scan(&n)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if isValidNumber(n) {
			break
		}

		fmt.Println("Please enter a 4-digit number.")
	}

	for result != 6174 {
		count++
		digits := []int{n / 1000, (n / 100) % 10, (n / 10) % 10, n % 10}
		sort.Ints(digits)

		// sorting logic:
		//The largest digit (digits[3]) goes in the thousands place (×1000)
		//The second largest (digits[2]) goes in the hundreds place (×100)
		//The third largest (digits[1]) goes in the tens place (×10)
		//The smallest digit (digits[0]) goes in the ones place
		//largest := digits[3]*1000 + digits[2]*100 + digits[1]*10 + digits[0]
		//smallest := digits[0]*1000 + digits[1]*100 + digits[2]*10 + digits[3]
		//revers for smallest

		// EXAMPLE if n=3524
		//digits = [3,5,2,4]
		//After sorting: digits = [2,3,4,5]
		//largest = 5×1000 + 4×100 + 3×10 + 2 = 5432
		//smallest = 2×1000 + 3×100 + 4×10 + 5 = 2345
		//result = 5432 - 2345 = 3087

		largest := digits[3]*1000 + digits[2]*100 + digits[1]*10 + digits[0]
		smallest := digits[0]*1000 + digits[1]*100 + digits[2]*10 + digits[3] // this shit took me frever to find out

		result = largest - smallest
		n = result

		fmt.Printf("%04d - %04d = %04d\n", largest, smallest, result)
	}

	fmt.Printf("\nKaprekar's constant reached in %d iterations.", count)
}

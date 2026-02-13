package main

import (
	"fmt"
)

func Min10(x []int) []int {
	newnumbers := []int{}
	for _, value := range x {
		if value >= 10 {
			newnumbers = append(newnumbers, value)
		}
	}
	return newnumbers
}

func main() {
	var count int
	var p int
	numbers := []int{}
	fmt.Printf("Введите кол-во чисел в слайсе\n")
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		fmt.Printf("Введите число №%d", i)
		fmt.Scan(&p)
		numbers = append(numbers, p)
	}

	new := Min10(numbers)
	fmt.Printf("\n Слайс с новым набором чисел")
	fmt.Println(new)
	fmt.Printf("\n Слайс со старым набором чисел")
	fmt.Println(numbers)
}

git config --global user.name "DS03-hash"
git config --global user.email "seredin11.03.2002@gmail.com"
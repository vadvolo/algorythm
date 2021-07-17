package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Fibonachy(n))
}

func Fibonachy(n int) (fibonachy int) {
	a := []int{0, 1}
	if n == 0 {
		return 0
	}
	for i := 2; i <= n; i++ {
		a = append(a, a[i-1] + a[i-2])
	}
	fibonachy = a[n]
	return
}

func f2() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)
}

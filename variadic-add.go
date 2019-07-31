package main

import "fmt"

func add(nums ...float64) (total float64) {
	for _, n := range nums {
		total += n
	}
	return
}
func main() {
	const pi = 3.14
	const r = 5
	const x = 30
	a := 12.56
	b := 40.00
	c := 10.25

	fmt.Println(add(pi, r))
	fmt.Println(add(r, x, a, b))
	fmt.Println(add(a, b, c, 2.2))
}

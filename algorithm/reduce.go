package main

func add(x, y int) int {
	return x + y
}

func Reduce(f func(int, int) int, a int, ary []int) int {
	for _, x := range ary {
		a = f(a, x)
	}
	return a
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	fmt.Println(Reduce(add, 1, a))
// }

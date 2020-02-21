package main

func square(x int) int {
	return x * x
}

func Mapping(f func(int) int, array []int) []int {
	buff := make([]int, len(array))
	for i, v := range array {
		buff[i] = f(v)
	}
	return buff
}

// func main() {
// 	a := []int{1, 2, 3}
// 	got := Mapping(square, a)
// 	fmt.Println(got)
// }

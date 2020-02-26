package filter

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 != 0
}

func Filter(f func(int) bool, ary []int) []int {
	buff := make([]int, 0)
	for _, v := range ary {
		if f(v) {
			buff = append(buff, v)
		}
	}
	return buff
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	got := Filter(isEven, a)
// 	fmt.Println(got)
// }

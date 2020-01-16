package main

func Quicksort(low, high int, buff []int) {
	pivot := buff[low+(high-low)/2]
	i, j := low, high
	for {
		for pivot > buff[i] {
			i++
		}
		for pivot < buff[j] {
			j--
		}
		if i >= j {
			break
		}
		buff[i], buff[j] = buff[j], buff[i]
		i++
		j--
	}
	if low < i-1 {
		Quicksort(low, i-1, buff)
	}
	if high > j+1 {
		Quicksort(j+1, high, buff)
	}
}

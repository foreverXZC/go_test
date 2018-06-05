package main

import "fmt"

var a = [10]int{5, 0, 5, 7, 0, 5, 7, 1, 0, -1}

func qsort(l, r int) {
	var i = l
	var j = r
	var x = a[(i+j)/2]
	for i < j {
		for a[i] < x {
			i++
		}
		for a[j] > x {
			j--
		}
		if i <= j {
			var t = a[i]
			a[i] = a[j]
			a[j] = t
			i++
			j--
		}
	}
	if l < j {
		qsort(l, j)
	}
	if i < r {
		qsort(i, r)
	}
}

func main() {
	qsort(0, 9)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", a[i])
	}
}

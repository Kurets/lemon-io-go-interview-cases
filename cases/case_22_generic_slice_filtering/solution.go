package case_22_generic_slice_filtering

import "fmt"

func Filter[T any](slice []T, pred func(T) bool) []T {
	var out []T
	for _, v := range slice {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	even := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println(even)
}

package main
import "fmt"
func sum(nums ...int) int {
	total := 0

	for _, val := range nums {
		total += val

	}
	return total
}
func main() {
	result := sum(1, 2, 3, 4, 5, 6, 7,8)
	fmt.Println("the sum is", result)
}

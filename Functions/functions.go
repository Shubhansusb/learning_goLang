package main

import "fmt"

func greet(name string) string {
	return "Hello Baby " + name
}

func add(a int, b int) int {
	return a + b
}

func Add_Prod(a, b int) (int, int) {
	return a + b, a * b
}

func sums(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v

	}
	return sum
}

type Rectangle struct {
	L, W float64
}

func (r Rectangle) Area() float64 {
	return r.L * r.W
}

func sum(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func getLang() (string, string, string) {
	return "golang", "javascript", "c++"
}
func main() {
	// fmt.Println(greet("shubhansu singh bhadoria"))
	// println(add(1, 4))
	// println(Add_Prod(10,5));

	// nums:=[]int{1,2,3,4,5}
	// println(sums(nums));

	// defer println("world");
	// println("hello");

	// rect := Rectangle{L: 10, W: 20}
	// println(rect.Area());

	// println(sum(1,2,3,4,5,6,7));
	a,b,c:=getLang();
	x,y,_:=getLang();
	fmt.Println(a,b,c)
	fmt.Println(x,y);

}

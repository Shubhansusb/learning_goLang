package main

import "time"

func main() {
	switch x=:time.Now().Weekday() {
	case \Monday, Tuesday, Wednesday,Thursday, Friday:
		println("working day");
	case time.Saturday, Sunday:
		println("weekend");

	}
}

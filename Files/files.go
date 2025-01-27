package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("namo.txt") // Open the file

	defer file.Close()

	if err != nil {
		panic(err)
	}

	buff := make([]byte, 10000)

	d, err := file.Read(buff)

	if err != nil {
		panic(err)
	}

	fmt.Println("data is ", string(buff), d)

	// if err != nil {
	// 	panic(err)
	// }
	// info, err := file.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File name is:", file.Name())
	// fmt.Println("is it a directory:", info.IsDir())
	// fmt.Println("mod time:", info.ModTime())
	// fmt.Println("size", info.Size())
	// fmt.Println("file mode", info.Mode())

}

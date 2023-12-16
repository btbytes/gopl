package main

import "fmt"
import "strings"
import "os"

func main() {
	fmt.Println("Name of the program", os.Args[0])
	fmt.Println("All the non program-name args: ", strings.Join(os.Args[1:], " "))
	for index, element := range os.Args[1:] {
		fmt.Println(index, element)
	}

}

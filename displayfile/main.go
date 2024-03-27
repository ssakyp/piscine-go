package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}
	name, error := ioutil.ReadFile(os.Args[1])
	if error != nil {
	}
	fmt.Print(string(name))
}

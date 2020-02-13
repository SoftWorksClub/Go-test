// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var sep string
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(sep + os.Args[i])
		sep = " "
	}
}
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for i, v := range os.Args {
			fmt.Printf("第%d个参数为%q\n", i, v)
		}
	}
}

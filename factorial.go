package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
	var value int
	fmt.Scanf("%d",&value)
	fmt.Println(fact(value))
}


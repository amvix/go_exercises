package main

import "fmt"

func main(){
	var instring string
	fmt.Scanf("%s", &instring)
	inarray := []byte(instring)
	for cont:=0; cont<len(inarray)-cont-1; cont++ {
		if  inarray[cont] != inarray[len(inarray)-cont-1] {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
	return
}

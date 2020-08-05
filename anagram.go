package main
import "fmt"

func main() {
	var tests int
	var str1 []byte
	var str2 []byte
	fmt.Scanf("%d", &tests)
	for n:=0; n < tests;  n++{
		fmt.Scanf("%s", &str1)
		fmt.Scanf("%s", &str2)

		deletes := 0
		for n:=0; n < len(str1); n++{
			found := false
			for nn:=0; nn < len(str2); nn++{
				if str1[n] == str2[nn] {
					str2[nn] = 0
					found = true
					break
				}
			}
			if !found {
				deletes++
				found = false
			}
		}
		for n:=0; n < len(str2); n++{
			if str2[n] != 0 {
				deletes++
			}
		}
		fmt.Println(deletes)
	}
}

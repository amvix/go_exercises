package main

import "fmt"

func main(){
   var testcases int
   fmt.Scanf("%d", &testcases)

   for cont:=0; cont < testcases; cont++ {
	   var seatnum int
	   var fpos int
	   fmt.Scanf("%d",&seatnum)
	   pos := seatnum%12

	   if pos == 0 {
		pos = 12
		fpos = 1
           }else{
		   fpos = ((12-pos)+1)
	   }
	   fseatn := seatnum+(fpos-pos)

	   fseat := ""
	   switch Abs(fpos-pos){
	     case 11: fseat = "WS" 
	     case 1: fseat = "WS" 
             case 3: fseat = "MS"
             case 9: fseat = "MS"
	     case 5: fseat = "AS"
	     case 7: fseat = "AS"
	   }
	   fmt.Println(fseatn,fseat)
   }
}

func Abs(value int) int{
  if value < 0{
	  return  -value
  }
  return value
}


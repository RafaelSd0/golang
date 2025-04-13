package main

import "fmt"

func main(){
	fmt.Println(Fatorial(4))
}


func Fatorial(n int) int {
	f  := 1
	for i := 1; i <= n; i++{
		f = f * i
	}
  return f
}



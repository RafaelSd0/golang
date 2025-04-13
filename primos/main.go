package main

import "fmt"

func main(){
	var n int
	n = 7
	if primo(n){
		fmt.Println("eh primo")
	} else {
		fmt.Println("não eh primo")
	}
	
}


func primo(n int) bool {
	if n <= 1{
		return false
	} else if n == 2{
		return true
	}
	
	for i := 2; i < n; i++{
		if n % i == 0 {
			return false
		}
	}
  return true
}



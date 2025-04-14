package main

import "fmt"

func main(){
	arr := [5]int{2,3,4,5,6}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		resto := arr[i] % 2
		if resto == 0  {
			fmt.Println(arr[i], " é par")
		} else {
			fmt.Println(arr[i], " é impar")
		}
	}
}

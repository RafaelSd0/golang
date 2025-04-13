package main

import "fmt"

func main(){
	fmt.Println(Hello("Rafael"))
}

// Olá retorna uma saudação para a pessoa nomeada.
func Hello(name string) string {
  // Retorna uma saudação que incorpora o nome em uma mensagem.
  message := fmt.Sprintf("Hi, %v. Welcome!", name) // ou var message string
                                                     // message = fmt.Sprintf("Hi, %v. Welcome!", name)
  return message
}



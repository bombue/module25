package main

import (
	"fmt"
	"log"
)

func main() {
	var someData string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&someData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы данные: %s\n", someData)
}

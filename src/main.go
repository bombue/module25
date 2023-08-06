package main

import (
	"fmt"
	"log"
)

func main() {
	var someData string
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&someData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %s\n", someData)
}

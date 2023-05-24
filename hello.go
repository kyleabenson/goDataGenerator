package main

import "fmt"
import "github.com/brianvoe/gofakeit/v6"

func main() {
    fmt.Println("Hello, World!")

	fmt.Println(gofakeit.Phone() )
	fmt.Println(gofakeit.CelebrityActor() )
}
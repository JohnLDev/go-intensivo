package main

import (
	"fmt"
	"time"
)

// ? pointer example
// import "fmt"

// type Car struct {
// 	Brand string
// }

// func main() {
// 	car := &Car{Brand: "Fiat"}

// 	copiaCarro := car
// 	copiaCarro.Brand = "Ford"
// 	fmt.Println(car.Brand)
// }
// ? ==========================

func task(name string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- name + " " + fmt.Sprint(i)
		time.Sleep(time.Second)
	}
}

// * thread 1
func main() {
	channel := make(chan string)

	// * go routine thread 2
	go func() {
		go task("A", channel)
		go task("B", channel)
		channel <- "Veio da t2"
	}()

	for msg := range channel {

		fmt.Println(msg)
	}

}

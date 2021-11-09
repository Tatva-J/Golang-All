package main

import (
	"fmt"
	"time"
)

// func main() {
// 	c := make(chan string)
// 	go count("sheep", c)
// 	for msg := range c {
// 		fmt.Println(msg)
// 	}
// 	// go count("fish")
// 	// // time.Sleep(time.Second * 5)
// 	// fmt.Scanln()

// }
// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }
// func main() {
// 	c := make(chan string, 2)
// 	c <- "hello"
// 	msg := <-c
// 	fmt.Println(msg)
// }
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()
	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}
	// }
	msg1 := <-c1
	msg2 := <-c2
	fmt.Println(msg1)
	fmt.Println(msg2)

}

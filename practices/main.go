package main

import (
	"fmt"
	"time"
)

/*
sent 1
sent 2
*/

func main() {
	ch := make(chan int, 2) // [1][2]

	ch <- 1
	fmt.Println("Sent 1")

	ch <- 2
	time.Sleep(2 * time.Second)
	fmt.Println("Sent 2")

	fmt.Println(<-ch, "dd")
	fmt.Println(<-ch, "dd")
	ch <- 3
	ch <- 4
	// ch <- 4

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

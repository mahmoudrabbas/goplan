package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	go n2(done)
	go n1(done)
	go n3(done)
	go n1(done)
	go n1(done)

	for range done {
		// <-done
	}

}

func n1(doneChannel chan bool) {
	fmt.Println("Hello Func 1")
	doneChannel <- true

}

func n2(doneChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello Func 2")

	doneChannel <- true

	close(doneChannel)
}

func n3(doneChannel chan bool) {
	fmt.Println("Hello Func 3")
	doneChannel <- true

}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	dones := make([]chan bool, 5)

// 	// done := make(chan bool)

// 	dones[0] = make(chan bool)
// 	go n2(dones[0])

// 	dones[1] = make(chan bool)
// 	go n1(dones[1])

// 	dones[2] = make(chan bool)
// 	go n3(dones[2])

// 	dones[3] = make(chan bool)
// 	go n1(dones[3])

// 	dones[4] = make(chan bool)
// 	go n1(dones[4])

// 	<-dones[0]
// 	<-dones[1]
// 	<-dones[2]
// 	<-dones[3]
// 	<-dones[4]

// }

// func n1(doneChannel chan bool) {
// 	fmt.Println("Hello Func 1")
// 	doneChannel <- true

// }

// func n2(doneChannel chan bool) {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Hello Func 2")

// 	doneChannel <- true
// }

// func n3(doneChannel chan bool) {
// 	fmt.Println("Hello Func 3")
// 	doneChannel <- true

// }

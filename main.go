package main

import (
	"fmt"
	"time"
)

func main() {
	tstagent()
}

func tstagent() {
	stop := make(chan bool)
	ready := make(chan error)
	fmt.Println("starting")
	go play(stop, ready)
	select {
	case <-ready:
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	err := check(1)
	fmt.Println(err)
	time.Sleep(1 * time.Second)
	err = check(2)
	fmt.Println(err)
	time.Sleep(1 * time.Second)

	stop <- true
	time.Sleep(5 * time.Second)
	err = check(3)
	fmt.Println(err)
	// fmt.Println(err)
	time.Sleep(8 * time.Second)
}

// func tstchn() {
// 	s := stp{
// 		Stop: make(chan bool),
// 	}
// 	go s.startEchoServer()
// 	time.Sleep(1 * time.Second)
// 	client()
// 	client()
// 	s.stopServer()
// 	client()
// 	time.Sleep(3 * time.Second)
// }

// func anotherfunction() {
// 	c := make(chan bool)
// 	play(c)
// 	return
// }

// func main() {
// 	// client()
// 	go anotherfunction()

// 	fmt.Println("aa")
// 	time.Sleep(8 * time.Second)
// }

package main

// import (
// 	"log"
// 	"net"
// 	"time"
// )

// func client() {
// 	c, err := net.Dial("unix", "/tmp/echo.sock")
// 	if err != nil {
// 		log.Fatal("Dial error", err)
// 	}
// 	defer c.Close()

// 	msg := "hi"
// 	_, err = c.Write([]byte(msg))
// 	if err != nil {
// 		log.Fatal("Write error:", err)
// 	}
// 	println("Client sent:", msg)
// 	time.Sleep(1e9)
// }

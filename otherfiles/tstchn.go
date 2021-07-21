package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// type stp struct {
// 	Stop chan bool
// }

// func echoServer(c net.Conn) {
// 	for {
// 		buf := make([]byte, 512)
// 		nr, err := c.Read(buf)
// 		if err != nil {
// 			return
// 		}

// 		data := buf[0:nr]
// 		println("Server got:", string(data))
// 		_, err = c.Write(data)
// 		if err != nil {
// 			log.Fatal("Write: ", err)
// 		}
// 	}
// }

// func (s *stp) startEchoServer() {
// 	l, err := net.Listen("unix", "/tmp/echo.sock")
// 	if err != nil {
// 		log.Fatal("listen error:", err)
// 	}

// 	for {
// 		select {
// 		case <-s.Stop:
// 			fmt.Println("stopping")
// 			return
// 		default:
// 			c, err := l.Accept()
// 			if err != nil {
// 				log.Fatal("accept error:", err)
// 			}

// 			buf := make([]byte, 512)
// 			nr, err := c.Read(buf)
// 			if err != nil {
// 				return
// 			}

// 			data := buf[0:nr]
// 			println("Server got:", string(data))
// 			_, err = c.Write(data)
// 			if err != nil {
// 				log.Fatal("Write: ", err)
// 			}
// 		}
// 	}
// }

// func client1(r io.Reader) {
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := r.Read(buf[:])
// 		if err != nil {
// 			return
// 		}
// 		println("Client got:", string(buf[0:n]))
// 	}
// }

// func client() {
// 	c, err := net.Dial("unix", "/tmp/go.sock")
// 	if err != nil {
// 		log.Fatal("Dial error", err)
// 	}
// 	defer c.Close()

// 	go client1(c)
// 	for {
// 		msg := "hi"
// 		_, err := c.Write([]byte(msg))
// 		if err != nil {
// 			log.Fatal("Write error:", err)
// 			break
// 		}
// 		println("Client sent:", msg)
// 		time.Sleep(1e9)
// 	}
// }

import (
	"fmt"
	"log"
	"net"
	"time"
)

type stp struct {
	Stop chan bool
}

func client() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()

	msg := "hi"
	_, err = c.Write([]byte(msg))
	if err != nil {
		log.Fatal("Write error:", err)
	}
	println("Client sent:", msg)
	time.Sleep(1e9)
}

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func (s *stp) startEchoServer() {
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	for {
		// time.Sleep(10000)
		// fmt.Println("aasdasdaa")
		// c, err := l.Accept()
		select {
		case <-s.Stop:
			fmt.Println("stopping")
			return
		default:
			c, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}
			go echoServer(c)
		}
	}
}

func (s *stp) stopServer() {
	fmt.Println("called stopping!")
	s.Stop <- true
}

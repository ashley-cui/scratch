package main

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	"time"

// 	"golang.org/x/crypto/ssh/agent"
// )

// func startserver(ctx context.Context) error {
// 	conn, err := net.Dial("unix", "/run/user/1000/keyring/ssh")
// 	if err != nil {
// 		return err
// 	}
// 	ac := agent.NewClient(conn)
// 	sock, err := net.Listen("unix", "/tmp/aaa")
// 	if err != nil {
// 		return err
// 	}
// 	go func() {
// 		<-ctx.Done()
// 		sock.Close()
// 	}()
// 	for {
// 		c, err := sock.Accept()
// 		if err != nil {
// 			if ctx.Err() == context.Canceled {
// 				return nil
// 			}
// 			return err

// 		}
// 		go agent.ServeAgent(ac, c)
// 	}
// }

// func check(i int) error {
// 	fmt.Println(i)
// 	conn, err := net.Dial("unix", "/tmp/aaa")
// 	if err != nil {
// 		fmt.Println("hehe")
// 		return err
// 	}
// 	ac := agent.NewClient(conn)
// 	fmt.Println(ac.List())
// 	return nil
// }

// func tstcontext() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go startserver(ctx)
// 	time.Sleep(time.Second)
// 	check(1)
// 	check(1)
// 	cancel()
// 	fmt.Println("afterclose")
// 	time.Sleep(3 * time.Second)
// 	// check(2)
// 	// check(2)
// 	// check(2)
// 	fmt.Println("sleep section")
// 	time.Sleep(8 * time.Second)
// }

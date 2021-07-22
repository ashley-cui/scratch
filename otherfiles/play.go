package main

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh/agent"
)

func play(stop chan bool, ready chan error) error {
	fmt.Println("here")
	conn, err := net.Dial("unix", "/run/user/1000/keyring/ssh")
	if err != nil {
		return err
	}
	ac := agent.NewClient(conn)
	sock, err := net.Listen("unix", "/tmp/aaa")
	if err != nil {
		fmt.Println(err)
		return err
	}
	ready <- nil
	for {
		select {
		case <-stop:
			fmt.Println("stopping")
			return nil
		default:
			fmt.Println("serving")
			c, _ := sock.Accept()
			go agent.ServeAgent(ac, c)
		}
	}
}

func check(i int) error {
	fmt.Println(i)
	conn, err := net.Dial("unix", "/tmp/aaa")
	if err != nil {
		fmt.Println("hehe")
		return err
	}
	ac := agent.NewClient(conn)
	fmt.Println(ac.List())
	return nil
}

// // struct with agent and channel
// // endagent -> send stop() on channel and remove sock
// type readOnlyAgent struct {
// 	agent.ExtendedAgent
// }

// func (a *readOnlyAgent) Add(_ agent.AddedKey) error {
// 	return errors.Errorf("adding new keys not allowed")
// }

// func (a *readOnlyAgent) Remove(_ ssh.PublicKey) error {
// 	return errors.Errorf("removing keys not allowed")
// }

// func (a *readOnlyAgent) RemoveAll() error {
// 	return errors.Errorf("removing keys not allowed")
// }

// func (a *readOnlyAgent) Lock(_ []byte) error {
// 	return errors.Errorf("locking agent not allowed")
// }

// func (a *readOnlyAgent) Extension(_ string, _ []byte) ([]byte, error) {
// 	return nil, errors.Errorf("extensions not allowed")
// }

// type AgentConfig struct {
// 	ID    string
// 	Paths []string
// }

// func makeAgent(config AgentConfig) agent.Agent {
// 	return &readOnlyAgent{}
// }
// func getSSHMounts()

// func play(stop chan bool, ready chan error) error {
// 	fmt.Println("here")
// 	conn, err := net.Dial("unix", "/run/user/1000/keyring/ssh")
// 	if err != nil {
// 		return err
// 	}
// 	ac := agent.NewClient(conn)
// 	sock, err := net.Listen("unix", "/tmp/aaa")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	ready <- nil
// 	go func() {
// 		for {
// 			fmt.Println("serving")
// 			c, _ := sock.Accept()
// 			go agent.ServeAgent(ac, c)
// 		}
// 	}()
// 	for {
// 		select {
// 		case <-stop:
// 			fmt.Println("stopping")
// 			return nil
// 		}
// 	}
// }

// func play2(a agent.Agent, l net.Listener) error {
// 	for {
// 		c, err := l.Accept()
// 		if err != nil {
// 			return err
// 		}

// 		go agent.ServeAgent(a, c)
// 	}
// }
// func play3() error {
// 	conn, err := net.Dial("unix", "/run/user/1000/keyring/ssh")
// 	if err != nil {
// 		return err
// 	}
// 	ac := agent.NewClient(conn)
// 	sock, err := net.Listen("unix", "/tmp/aaa")
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func sshAgentFwd(stop chan bool) {

// }

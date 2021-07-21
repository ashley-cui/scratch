package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net"
// 	"os"
// 	"path/filepath"

// 	"golang.org/x/crypto/ssh"
// 	"golang.org/x/crypto/ssh/agent"
// )

// type SSHProvider interface {
// 	StartAgent() error
// 	StopAgent() error
// }
// type SSHSocketFwd struct {
// 	SourcePath string
// 	ServePath  string
// 	Stop       chan bool
// }
// type SSHKeyFwd struct {
// 	Keys      []interface{}
// 	ServePath string
// 	Stop      chan bool
// }

// func NewSocketProvider(sourcePath string) *SSHSocketFwd {
// 	serveDir, err := ioutil.TempDir("", ".buildah-ssh-sock")
// 	if err != nil {
// 		fmt.Println("ERROR")
// 	}
// 	servePath := filepath.Join(serveDir, "ssh_auth_sock")
// 	if err != nil {
// 		fmt.Println("ERROR")
// 	}
// 	return &SSHSocketFwd{
// 		SourcePath: sourcePath,
// 		ServePath:  servePath,
// 		Stop:       make(chan bool),
// 	}
// }
// func NewKeyProvider(keys []interface{}) *SSHKeyFwd {
// 	serveDir, err := ioutil.TempDir("", ".buildah-ssh-sock")
// 	if err != nil {
// 		fmt.Println("ERROR")
// 	}
// 	servePath := filepath.Join(serveDir, "ssh_auth_sock")
// 	if err != nil {
// 		fmt.Println("ERROR")
// 	}
// 	return &SSHKeyFwd{
// 		Keys:      keys,
// 		ServePath: servePath,
// 		Stop:      make(chan bool),
// 	}
// }

// func NewSSHProvider(paths []string) SSHProvider {
// 	var keys []interface{}
// 	var socket string
// 	if len(paths) == 0 {
// 		socket := os.Getenv("SSH_AUTH_SOCK")
// 		if socket == "" {
// 			fmt.Println("ERROR")
// 		}
// 	}
// 	for _, p := range paths {
// 		if socket != "" {
// 			fmt.Println("ERROR")
// 		}

// 		fi, err := os.Stat(p)
// 		if err != nil {
// 			fmt.Println("ERROR")
// 		}
// 		if fi.Mode()&os.ModeSocket > 0 {
// 			if len(keys) == 0 {
// 				socket = p
// 			} else {
// 				fmt.Println("ERROR")
// 			}
// 			continue
// 		}

// 		f, err := os.Open(p)
// 		if err != nil {
// 			fmt.Println("ERROR failed to open ")
// 		}
// 		dt, err := ioutil.ReadAll(&io.LimitedReader{R: f, N: 100 * 1024})
// 		if err != nil {
// 			fmt.Println("ERROR failed to read ")
// 		}

// 		k, err := ssh.ParseRawPrivateKey(dt)
// 		if err != nil {
// 			keys = append(keys, k)
// 		}
// 	}
// 	if socket != "" {
// 		return NewSocketProvider(socket)
// 	}
// 	return NewKeyProvider(keys)
// }

// func (s *SSHKeyFwd) StartAgent() error {
// 	a := agent.NewKeyring()
// 	for _, k := range s.Keys {
// 		if err := a.Add(agent.AddedKey{PrivateKey: k}); err != nil {
// 			fmt.Println("ERROR failed to add to agent")
// 		}
// 	}
// 	sock, err := net.Listen("unix", s.ServePath)
// 	if err != nil {
// 		fmt.Println("aaaaaaa")
// 	}
// 	for {
// 		select {
// 		case <-s.Stop:
// 			fmt.Println("stopping")
// 			return nil
// 		default:
// 			c, _ := sock.Accept()
// 			err := agent.ServeAgent(a, c)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// }
// func (s *SSHKeyFwd) StopAgent() error {
// 	s.Stop <- true
// 	close(s.Stop)
// 	return nil
// }

// func (s *SSHSocketFwd) StartAgent() error {
// 	conn, err := net.Dial("unix", s.SourcePath)
// 	if err != nil {
// 		fmt.Println("aaaaaaa")
// 	}
// 	defer conn.Close()
// 	ac := agent.NewClient(conn)
// 	sock, err := net.Listen("unix", s.ServePath)
// 	if err != nil {
// 		fmt.Println("aaaaaaa")
// 	}
// 	for {
// 		select {
// 		case <-s.Stop:
// 			fmt.Println("stopping")
// 			return nil
// 		default:
// 			c, _ := sock.Accept()
// 			err := agent.ServeAgent(ac, c)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// }

// func (s *SSHSocketFwd) StopAgent() error {
// 	s.Stop <- true
// 	os.RemoveAll(s.ServePath)
// 	return nil
// }

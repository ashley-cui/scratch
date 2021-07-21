package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net"
// 	"os"

// 	"golang.org/x/crypto/ssh"
// 	"golang.org/x/crypto/ssh/agent"
// )

// // func newAgent(paths []string) (agent.Agent, error) {
// // 	if len(paths) == 0 || len(paths) == 1 && paths[0] == "" {
// // 		paths = []string{os.Getenv("SSH_AUTH_SOCK")}
// // 	}
// // 	var socket bool
// // 	a := agent.NewKeyring()
// // 	for _, p := range paths {
// // 		if socket != nil {
// // 			return source{}, errors.New("only single socket allowed")
// // 		}
// // 		fi, err := os.Stat(p)
// // 		if err != nil {
// // 			return err
// // 		}
// // 		if fi.Mode()&os.ModeSocket > 0 {
// // 			socket = true
// // 			continue

// // 		}
// // 	}
// // }
// func play(stop chan bool, agent agent.Agent) {
// 	fmt.Println("here")
// 	conn, _ := net.Dial("unix", "/run/user/1000/keyring/ssh")
// 	ac := agent.NewClient(conn)
// 	sock, err := net.Listen("unix", "/tmp/aaa")
// 	if err != nil {
// 		fmt.Println("aaaaaaa")
// 	}
// 	for {
// 		select {
// 		case <-stop:
// 			fmt.Println("stopping")
// 		default:
// 			c, _ := sock.Accept()
// 			agent.ServeAgent(ac, c)
// 		}
// 	}
// }

// type sshSrc interface {
// 	forwardAgent(stop chan bool)
// }
// type socketFwd struct {
// 	path      string
// 	stop      chan bool
// 	agentPath string
// }
// func serveAgents(paths []string){

// 	fi, err := os.Stat(p)
// 	if err != nil {
// 		return source{}, errors.WithStack(err)
// 	}
// 	if fi.Mode()&os.ModeSocket > 0 {
// 		socket = &socketDialer{path: p, dialer: unixSocketDialer}
// 		continue
// 	}
// }
// func serveSocket(stop chan bool, socketpath string, newagentpath string) {
// 	conn, _ := net.Dial("unix", socketpath)
// 	ac := agent.NewClient(conn)
// 	sock, err := net.Listen("unix", newagentpath)
// 	if err != nil {
// 		fmt.Println("aaaaaaa")
// 	}
// 	for {
// 		select {
// 		case <-stop:
// 			fmt.Println("stopping")
// 		default:
// 			c, _ := sock.Accept()
// 			agent.ServeAgent(ac, c)
// 		}
// 	}
// }
// func serveKeys(stop chan bool, keyfiles []string, newagentpath string) {
// 	a := agent.NewKeyring()
// 	for _, p := range keyfiles {
// 		f, err := os.Open(p)
// 		if err != nil {
// 			fmt.Println("stopping")
// 		}
// 		dt, err := ioutil.ReadAll(&io.LimitedReader{R: f, N: 100 * 1024})
// 		if err != nil {
// 			fmt.Println("stopping")
// 		}
// 		k, err := ssh.ParseRawPrivateKey(dt)
// 		if err := a.Add(agent.AddedKey{PrivateKey: k}); err != nil {
// 			fmt.Println("stopping")
// 		}
// 	}
// 	sock, err := net.Listen("unix", newagentpath)
// 	if err != nil {
// 		fmt.Println("stopping")
// 	}
// 	for {
// 		select {
// 		case <-stop:
// 			fmt.Println("stopping")
// 		default:
// 			c, _ := sock.Accept()
// 			agent.ServeAgent(a, c)
// 		}
// 	}
// }

// // func sshAgentFwd(stop chan bool) {

// // }

// // // struct with agent and channel
// // // endagent -> send stop() on channel and remove sock
// // type readOnlyAgent struct {
// // 	agent.ExtendedAgent
// // }

// // func (a *readOnlyAgent) Add(_ agent.AddedKey) error {
// // 	return errors.Errorf("adding new keys not allowed")
// // }

// // func (a *readOnlyAgent) Remove(_ ssh.PublicKey) error {
// // 	return errors.Errorf("removing keys not allowed")
// // }

// // func (a *readOnlyAgent) RemoveAll() error {
// // 	return errors.Errorf("removing keys not allowed")
// // }

// // func (a *readOnlyAgent) Lock(_ []byte) error {
// // 	return errors.Errorf("locking agent not allowed")
// // }

// // func (a *readOnlyAgent) Extension(_ string, _ []byte) ([]byte, error) {
// // 	return nil, errors.Errorf("extensions not allowed")
// // }

// // type AgentConfig struct {
// // 	ID    string
// // 	Paths []string
// // }

// // func makeAgent(config AgentConfig) agent.Agent {
// // 	return &readOnlyAgent{}
// // }
// // func getSSHMounts()

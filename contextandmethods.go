package main

import (
	"context"
	"time"

	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type Source interface {
	NewAgent() (*AgentServer, error)
}
type SSHSocketFwd struct {
	SourcePath string
	ServePath  string
}
type SSHKeyFwd struct {
	Keys      []interface{}
	ServePath string
}

type AgentServer struct {
	Agent  agent.Agent
	Cancel func()
}

func NewSocketSource(sourcePath string) (*SSHSocketFwd, error) {
	return &SSHSocketFwd{
		SourcePath: sourcePath,
	}, nil
}
func NewKeySource(keys []interface{}) (*SSHKeyFwd, error) {
	return &SSHKeyFwd{
		Keys: keys,
	}, nil
}

func NewSource(paths []string) (Source, error) {
	var keys []interface{}
	var socket string
	if len(paths) == 0 {
		socket := os.Getenv("SSH_AUTH_SOCK")
		if socket == "" {
			return nil, errors.New("$SSH_AUTH_SOCK not set")
		}
	}
	for _, p := range paths {
		if socket != "" {
			return nil, errors.New("only one socket is allowed")
		}

		fi, err := os.Stat(p)
		if err != nil {
			return nil, err
		}
		if fi.Mode()&os.ModeSocket > 0 {
			if len(keys) == 0 {
				socket = p
			} else {
				return nil, errors.New("cannot mix keys and socket file")
			}
			continue
		}

		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		dt, err := ioutil.ReadAll(&io.LimitedReader{R: f, N: 100 * 1024})
		if err != nil {
			return nil, err
		}

		k, err := ssh.ParseRawPrivateKey(dt)
		if err != nil {
			keys = append(keys, k)
		}
	}
	if socket != "" {
		return NewSocketSource(socket)
	}
	return NewKeySource(keys)
}

func (s *SSHKeyFwd) NewAgent() (*AgentServer, error) {
	a := agent.NewKeyring()
	for _, k := range s.Keys {
		if err := a.Add(agent.AddedKey{PrivateKey: k}); err != nil {
			return nil, errors.Wrap(err, "failed to create ssh agent")
		}
	}

	return &AgentServer{
		Agent: a,
	}, nil
}

func (s *SSHSocketFwd) NewAgent() (*AgentServer, error) {
	conn, err := net.Dial("unix", s.SourcePath)
	if err != nil {
		return nil, err
	}
	ac := agent.NewClient(conn)
	// defer conn.Close()
	return &AgentServer{
		Agent: ac,
	}, nil
}

func (a *AgentServer) ServeAgent(ctx context.Context) error { // this function also probably needs waitgroup
	// serveDir, err := ioutil.TempDir("", ".buildah-ssh-sock")
	// if err != nil {
	// 	return err
	// }
	// servePath := filepath.Join(serveDir, "ssh_auth_sock")
	servePath := "/tmp/aaa"

	sock, err := net.Listen("unix", servePath)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		sock.Close()
	}()
	for {
		fmt.Println("serving!!")
		c, err := sock.Accept()
		if err != nil {
			if ctx.Err() == context.Canceled {
				return nil
			}
			return err
		}
		go agent.ServeAgent(a.Agent, c)
	}
}

////////////////////// some functions to test stuff /////////////////////////

func drive() (context.CancelFunc, error) {
	path := os.Getenv("SSH_AUTH_SOCK") //test stuff, this shou8d already be in buildah

	// waitGroup := sync.WaitGroup{}
	// this part is in a for loop in  buildah for every --mount flag
	source, err := NewSource([]string{path})
	if err != nil {
		return nil, err
	}
	ag, err := source.NewAgent()
	// listofagents := append(listofagents, ag)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	ag.Cancel = cancel
	// waitGroup.Add(1)
	// go ag.ServeAgent(ctx, waitgroup)
	go ag.ServeAgent(ctx)
	// ^^ end of for loop

	// buildah, some work is done here
	// i guess instead of passing around ctx, we can just  pass around our list of agents
	// so we can call agent.cancel() later on
	// in buildah, there's a defer cleanuprunmounts(agnet, waitgroup) here that should clean up agents

	// waitGroup.Wait()
	return cancel, nil
}

// func pretnedbuildah(sources map[string]Source, sshids []string) {
// 	var agents []AgentServer
// 	for _, id := range sshids {
// 		if src, ok := sources[id]; !ok {
// 			ag, err := src.NewAgent()
// 			agents.Appe
// 		}
// 	}
// }

// func cleanuprunmounts(agents, waitgroup ,...){
// 	for agent:- range agent{
// 		agent.Cancel()
// 		might also need to cleanup some other connections?
//  ...
// 	}

// }
func m() {
	cancel, _ := drive()
	time.Sleep(5 * time.Second)
	aaaa(1)
	aaaa(1)
	cancel()
	time.Sleep(2 * time.Second)
	aaaa(1)

}

func aaaa(i int) error {
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

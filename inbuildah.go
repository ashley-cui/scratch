package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type Source interface {
	StartAgent(stop chan bool) error
	StopAgent(stop chan bool) error
}
type SSHSocketFwd struct {
	SourcePath string
	ServePath  string
}
type SSHKeyFwd struct {
	Keys      []interface{}
	ServePath string
}

func NewSocketSource(sourcePath string) (*SSHSocketFwd, error) {
	serveDir, err := ioutil.TempDir("", ".buildah-ssh-sock")
	if err != nil {
		return nil, err
	}
	servePath := filepath.Join(serveDir, "ssh_auth_sock")
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println("creating socket src")
	return &SSHSocketFwd{
		SourcePath: sourcePath,
		ServePath:  servePath,
	}, nil
}
func NewKeySource(keys []interface{}) (*SSHKeyFwd, error) {
	serveDir, err := ioutil.TempDir("", ".buildah-ssh-sock")
	if err != nil {
		return nil, err
	}
	servePath := filepath.Join(serveDir, "ssh_auth_sock")
	return &SSHKeyFwd{
		Keys:      keys,
		ServePath: servePath,
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
		fmt.Println("got socket!")
		return NewSocketSource(socket)
	}
	fmt.Println("got keys!")
	return NewKeySource(keys)
}

func (s *SSHKeyFwd) StartAgent(stop chan bool) error {
	a := agent.NewKeyring()
	for _, k := range s.Keys {
		if err := a.Add(agent.AddedKey{PrivateKey: k}); err != nil {
			fmt.Println("ERROR failed to add to agent")
		}
	}
	sock, err := net.Listen("unix", s.ServePath)
	if err != nil {
		fmt.Println("aaaaaaa")
		return err
	}
	for {
		select {
		case <-stop:
			fmt.Println("stopping")
			return nil
		default:
			c, _ := sock.Accept()
			err := agent.ServeAgent(a, c)
			if err != nil {
				return err
			}
		}
	}
}
func (s *SSHKeyFwd) StopAgent(stop chan bool) error {
	stop <- true
	close(stop)
	return nil
}

func (s *SSHSocketFwd) StartAgent(stop chan bool) error {
	conn, err := net.Dial("unix", s.SourcePath)
	if err != nil {
		fmt.Println("aaaaaaa")
		return err
	}
	defer conn.Close()
	ac := agent.NewClient(conn)
	sock, err := net.Listen("unix", s.ServePath)
	if err != nil {
		fmt.Println("aaaaaaa")
		return err
	}
	for {
		select {
		case <-stop:
			fmt.Println("stopping")
			return nil
		default:
			c, _ := sock.Accept()
			err := agent.ServeAgent(ac, c)
			if err != nil {
				return err
			}
		}
	}
}

func (s *SSHSocketFwd) StopAgent(stop chan bool) error {
	stop <- true
	os.RemoveAll(s.ServePath)
	return nil
}

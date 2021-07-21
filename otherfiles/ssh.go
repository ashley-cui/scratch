package main

// type SSHAgent struct {
// 	agent      agent.Agent
// 	socketPath string
// 	listener   net.Listener
// 	stop       chan bool
// }

// func NewSSHAgent(socketPath string) (*SSHAgent, error) {
// 	a := &SSHAgent{
// 		agent:      agent.NewKeyring(),
// 		stop:       make(chan bool),
// 		socketPath: socketPath,
// 	}
// 	listener, err := net.Listen("unix", a.socketPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	a.listener = listener
// 	return a, nil
// }
// func (a *SSHAgent) runAgent(socketPath string, keys []) {
// 	quit := make(chan bool)
// 	go func(c io.ReadWriter) {
// 		err := agent.ServeAgent(a.agent, c)
// 		if err != nil {
// 			fmt.Printf("could not serve ssh agent %v", err)
// 		}
// 	}(c)

// 	fmt.Println("stopping!")
// }

// func (a SSHAgent) stopAgent() {
// 	a.stop
// }

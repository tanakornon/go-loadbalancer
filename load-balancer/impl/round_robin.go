package balancer

import (
	"fmt"
	"net/http"
	"sync"
)

type RoundRobin struct {
	servers  []string
	sequence int
	mu       sync.Mutex
}

func NewRoundRobin(servers []string) *RoundRobin {
	return &RoundRobin{
		servers:  servers,
		sequence: 0,
	}
}

func (lb *RoundRobin) GetNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	server := lb.servers[lb.sequence]
	lb.sequence = (lb.sequence + 1) % len(lb.servers)

	return server
}

func (lb *RoundRobin) Handler(w http.ResponseWriter, r *http.Request) {
	server := lb.GetNextServer()

	fmt.Fprintf(w, "Approach: Round Robin\n")
	fmt.Fprintf(w, "Request served by: %s\n", server)
}

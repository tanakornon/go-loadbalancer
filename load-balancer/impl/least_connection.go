package balancer

import (
	"fmt"
	"net/http"
	"sync"
)

type LeastConnection struct {
	servers          []string
	connectionCounts map[string]int
	mu               sync.Mutex
}

func NewLeastConnection(servers []string) *LeastConnection {
	lb := &LeastConnection{
		servers:          servers,
		connectionCounts: make(map[string]int),
	}
	for _, server := range servers {
		lb.connectionCounts[server] = 0
	}
	return lb
}

func (lb *LeastConnection) GetServerWithLeastConnections() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	var minServer string
	minConnections := -1

	for server, connections := range lb.connectionCounts {
		if minConnections == -1 || connections < minConnections {
			minConnections = connections
			minServer = server
		}
	}

	return minServer
}

func (lb *LeastConnection) Handler(w http.ResponseWriter, r *http.Request) {
	server := lb.GetServerWithLeastConnections()

	fmt.Fprintf(w, "Approach: Least Connection\n")
	fmt.Fprintf(w, "Request served by: %s\n", server)
	defer fmt.Fprintf(w, "Connection count: %+v\n", lb.connectionCounts)

	lb.mu.Lock()
	lb.connectionCounts[server]++
	lb.mu.Unlock()
}

package main

import (
	"fmt"
	balancer "load-balancer/impl"
	"net/http"
)

func main() {

	servers := []string{
		"Server1:8080",
		"Server2:8080",
		"Server3:8080",
	}

	// loadBalancer := balancer.NewRoundRobin(servers)
	loadBalancer := balancer.NewLeastConnection(servers)

	http.HandleFunc("/", loadBalancer.Handler)
	fmt.Println("Load balancer is running on :8080...")
	http.ListenAndServe(":8080", nil)
}

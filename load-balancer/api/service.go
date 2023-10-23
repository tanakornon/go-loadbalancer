package api

import "net/http"

type LoadBalancer interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

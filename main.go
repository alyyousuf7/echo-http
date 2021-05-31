package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

type HTTPHandler struct{}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// write the requested status code
	requestedStatus := r.Header.Get("requested-status")
	statusInt, err := strconv.ParseInt(requestedStatus, 10, 32)
	if err == nil {
		w.WriteHeader(int(statusInt))
	}

	// write the requested url
	fmt.Fprintf(w, "%s %s\n", r.Method, r.URL)
}

func main() {
	portPtr := flag.Int("port", 80, "listening port")
	flag.Parse()

	port := *portPtr
	addr := fmt.Sprintf(":%d", port)
	handler := HTTPHandler{}

	// Get all interfaces and print ipv4 address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				fmt.Printf("http://%s:%d\n", ipnet.IP, port)
			}
		}
	}

	// start the server
	if err := http.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}

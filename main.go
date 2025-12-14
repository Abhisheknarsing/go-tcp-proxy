package main

import (
	"log"
	"net"
	"go-tcp-proxy/config"
	"go-tcp-proxy/proxy"
)

func main() {
	cfg := config.Load()

	listener, err := net.Listen("tcp", cfg.LocalAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", cfg.LocalAddr, err)
	}

	log.Printf("Proxying %s -> %s", cfg.LocalAddr, cfg.RemoteAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go proxy.HandleConnection(conn, cfg.RemoteAddr)
	}
}

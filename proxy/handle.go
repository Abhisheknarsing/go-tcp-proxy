package proxy

import (
	"io"
	"log"
	"net"
)

func HandleConnection(local net.Conn, remoteAddr string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Printf("Failed to dial remote %s: %v", remoteAddr, err)
		local.Close()
		return
	}

	go copyStream(local, remote)
	go copyStream(remote, local)
}

func copyStream(dst io.WriteCloser, src io.ReadCloser) {
	defer dst.Close()
	defer src.Close()
	io.Copy(dst, src)
}

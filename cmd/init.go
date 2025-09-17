package main

import (
	"kvstore/server"
	"log"
	"net"
	"strings"
)

func runApp() {
	srv := server.NewServer("", 0)
	srv.Start(handleconn)
}

func handleconn(c net.Conn) {
	cmsg := make([]byte, 1024)
	for {
		readlen, err := c.Read(cmsg)
		if err != nil {
			log.Printf("msg read: %s", string(cmsg))
		}

		var resp string
		if strings.EqualFold(string(cmsg[:readlen]), "ping") {
			resp = "PONG"
		} else {
			resp = "invalid command,as of now only 'ping' is supported in this version :)"
		}

		if _, err := c.Write([]byte(resp)); err != nil {
			log.Fatalf("writing response to client %s failed: %s", c.RemoteAddr(), err.Error())
			continue
		}
	}
}

package server

import (
	"fmt"
	"log"
	"net"
)

type ServerInfo struct {
	IP     string
	Port   int16
	Server net.Listener
}

func NewServer(ip string, port int16) *ServerInfo {
	if ip == "" {
		ip = "127.0.0.1"
	}
	if port == 0 {
		port = 4455
	}
	addr := fmt.Sprintf("%s:%d", ip, port)

	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error while dialing tcp, addr:%s, %s", addr, err.Error())
		return nil
	}

	log.Printf("tcp server started on %s", addr)

	return &ServerInfo{
		IP:     ip,
		Port:   port,
		Server: server,
	}
}

func (si *ServerInfo) Start(handleconn func(net.Conn)) {
	for {
		conn, err := si.Server.Accept()
		if err != nil {
			log.Printf("error while accepting TCP connection on sever: %s", err.Error())
		}
		_ = conn

		log.Printf("client connected on server, ip:%s", conn.RemoteAddr().String())
		handleconn(conn)
	}
}

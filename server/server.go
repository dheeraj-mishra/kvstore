package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type ServerInfo struct {
	IP     string
	Port   int16
	Server net.Listener
}

func NewServer(ip string, port int16) (*ServerInfo, func(net.Conn)) {
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
		return nil, nil
	}

	log.Printf("tcp server started on %s", addr)

	return &ServerInfo{
		IP:     ip,
		Port:   port,
		Server: server,
	}, handleconn
}

func handleconn(c net.Conn) {
	defer c.Close()
	cmsg := make([]byte, 1024)
	for {
		readlen, err := c.Read(cmsg)
		if err == io.EOF {
			log.Printf("client connection closed: %s", c.RemoteAddr())
			break
		} else if err != nil {
			log.Printf("msg read: %s", string(cmsg))
		}

		var resp string
		if strings.EqualFold(string(cmsg[:readlen]), "ping") {
			resp = "PONG"
		} else {
			resp = "invalid command,as of now only 'ping' is supported in this version :)"
		}

		if _, err := c.Write([]byte(resp)); err != nil {
			log.Printf("writing response to client %s failed: %s", c.RemoteAddr(), err.Error())
			continue
		}

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

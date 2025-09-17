package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type ClientInfo struct {
	IP     string
	Port   int16
	Client net.Conn
}

func main() {
	var ci ClientInfo = ClientInfo{}
	// fmt.Println("len", len(os.Args))
	// fmt.Println("args", os.Args)

	if len(os.Args) == 1 {
		ci.IP = "127.0.0.1"
		ci.Port = 4455
	} else if len(os.Args) == 2 {
		ci.IP = os.Args[1]
		port, _ := strconv.Atoi(os.Args[2])
		ci.Port = int16(port)
	} else {
		log.Fatal("invalid args. valid: <ip> <port> OR no args for default")
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ci.IP, ci.Port))
	if err != nil {
		log.Fatalf("connecting to kvstore server %s:%d failed: %s", ci.IP, ci.Port, err.Error())
	}
	ci.Client = conn

	var cmd string
	ba := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		if _, err := ci.Client.Write([]byte(cmd)); err != nil {
			log.Fatalf("writing to sever failed, msg:%s, error: %s", cmd, err.Error())
			continue
		}

		readlen, err := ci.Client.Read(ba)
		if err != nil {
			log.Fatalf("reading from sever failed, error: %s", err.Error())
			continue
		}

		fmt.Println(string(ba[:readlen]))
	}
}

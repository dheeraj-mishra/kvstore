package main

import (
	"kvstore/server"
)

func runApp() {
	srv, handleconn := server.NewServer("", 0)
	srv.Start(handleconn)
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("arg err!")
		os.Exit(0)
	}
	port := os.Args[1]
	fmt.Println("Port is:", port)
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(addr)
	tcpListener, _ := net.ListenTCP("tcp4", addr)
	defer tcpListener.Close()
	for {
		client, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle_client(client)
	}
}

func handle_client(tcp *net.TCPConn) {
	var buff = make([]byte, 1024)
	defer tcp.Close()
	for {
		n, err := tcp.Read(buff)
		if n <= 0 || err != nil {
			break
		}
		fmt.Println(string(buff[:]))
		tcp.Write(buff)
	}
}

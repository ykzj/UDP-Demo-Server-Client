package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var port int
	var host string
	flag.StringVar(&host, "h", "", "host to connect")
	flag.IntVar(&port, "p", 10000, "port")
	flag.Parse()
	if host == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("udp", addr)
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}
	conn.Write([]byte("hello world!"))
	fmt.Println("send msg")
	var msg [20]byte
	conn.Read(msg[0:])
	fmt.Println("msg is:", string(msg[0:20]))
}

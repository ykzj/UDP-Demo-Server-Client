package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvUDPMsg(conn *net.UDPConn) {
	for {
		var buf [20]byte
		n, raddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return
		}
		fmt.Println("msg is:", string(buf[0:n]))

		_, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
		checkError(err)
	}

}

func main() {
	var port int
	var host string
	flag.IntVar(&port, "p", 10000, "port, default to 10000")
	flag.StringVar(&host, "h", "0.0.0.0", "host to connect")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", host, port)
	udp_addr, err := net.ResolveUDPAddr("udp", addr)
	checkError(err)

	conn, err := net.ListenUDP("udp", udp_addr)
	defer conn.Close()
	checkError(err)

	recvUDPMsg(conn)
}

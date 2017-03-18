package socket

import (
	"fmt"
	"net"
)

func Init() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8085")
	conn, _ := listener.Accept()
	fmt.Println(conn)
}

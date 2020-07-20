package socket

import (
	"fmt"
	"net"
)

func Write(data []byte, printer string) int { //Write
	fmt.Println("write call")
	conn, _ := net.Dial("tcp", printer)
	fmt.Println(data)
	byteSent, _ := conn.Write(data)
	//check error
	fmt.Println(byteSent)
	return byteSent
}

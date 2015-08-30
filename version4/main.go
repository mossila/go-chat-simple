package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func clientHandler(c net.Conn) {
	defer c.Close()
	for {
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Print(msg)
	}
}

/*ChatServer start chatserver on tcp with specific port
 * Because this is exported function
 * if it hos no comment here you will get warning like this
 * > exported function ChatServer should have comment or be unexported
 */
func ChatServer(port string) {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go clientHandler(conn)
	}
}

func main() {
	port := os.Args[1]
	ChatServer(port)
}

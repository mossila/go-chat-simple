package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
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

		go func(c net.Conn) {
			defer c.Close()
			for {
				msg, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				fmt.Print(msg)
			}
		}(conn)
	}
}

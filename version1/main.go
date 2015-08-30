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

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Print(msg)
		conn.Close()
	}
}

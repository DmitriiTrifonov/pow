package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {

		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(msg[:len(msg)-1])
	}
}

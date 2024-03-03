package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type client struct {
	conn net.Conn
	name string
}

var clients []client

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	newClient := client{conn: conn, name: name}
	clients = append(clients, newClient)

	fmt.Printf("%s has joined the chat\n", name)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		for _, c := range clients {
			c.conn.Write([]byte(fmt.Sprintf("%s: %s\n", name, message)))
		}
	}

	removeClient(newClient)
}

func removeClient(clientToRemove client) {
	for i, c := range clients {
		if c == clientToRemove {
			clients = append(clients[:i], clients[i+1:]...)
			fmt.Printf("%s has left the chat\n", clientToRemove.name)
			break
		}
	}
}

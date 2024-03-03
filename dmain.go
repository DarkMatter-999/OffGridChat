package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/DarkMatter-999/OffGridChat/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	HOST = "localhost"
	PORT = "6969"
)

type Host struct {
	IP   string
	Name string
}

var Hosts []Host

func main() {
	host := flag.String("host", HOST, "Server host address")
	port := flag.String("port", PORT, "Server port")

	flag.Parse()
	
	log.Printf("Using IP %s:%s", *host, *port);


	lis, err := net.Listen("tcp", *host + ":" + *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	var wg sync.WaitGroup

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var clientaddr string
		fmt.Printf("Enter Client address : ")
		fmt.Scanf("%s", &clientaddr)
		connectAndSendMessage(clientaddr)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh 

	grpcServer.Stop()

	lis.Close()

	wg.Wait()
}

func connectAndSendMessage(serverAddress string) {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := chat.NewChatServiceClient(conn)

	messageReq := &chat.MessageRequest{ RecipientIp: serverAddress, Message: "This is a test"}

	response, err := client.SendMessage(context.Background(), messageReq)
	if err != nil {
		log.Fatalf("SendMessage failed: %v", err)
	}

	log.Printf("SendMessage response: %v", response)
}


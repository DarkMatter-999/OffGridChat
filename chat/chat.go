package chat

import (
	"context"
	"log"
)

type Server struct {}
	

func (s *Server) Discover(context context.Context, message *DiscoverPacket) (*DiscoverPacket, error) {
	log.Printf("Recieved %s", message.Host);
	
	return &DiscoverPacket{Host: []*Host{{Ip: "localhost:6969", Name: "6969"},}}, nil
}

func (s *Server) SendMessage(context context.Context, message *MessageRequest) (*MessageResponse, error) {
	log.Printf("%s :  %s", message.RecipientIp, message.Message)
	return &MessageResponse{Message: "ok"}, nil	
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}

package main

import (
	contactpb "cloudrun-grpc/go-contact/proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var data = map[string]string{
	"555-110022": "Mike F",
	"555-220033": "Rob U",
	"555-330044": "Tanya C",
}

type ContactServer struct {
	contactpb.UnimplementedContactServer
}

func (s *ContactServer) GetContact(ctx context.Context, req *contactpb.ContactRequest) (*contactpb.ContactReply, error) {
	return &contactpb.ContactReply{
		Name: data[req.GetPhoneNumber()],
	}, nil
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("%v", err)
	}
	s := grpc.NewServer()
	contactpb.RegisterContactServer(s, &ContactServer{})
	log.Printf("ContactServer tcp/%v", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Fatalf("%v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "mini_project_dis_sys"
	"net"
	"sync"
)

const (
	port = ":50051"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServiceServer) Publish(ctx context.Context, in *pb.ChatMessage) (*pb.ChatMessage, error) {
	fmt.Printf("Received message: %v %v \n", in.GetSender(), in.GetMessage())
	msgChannel <- *in
	return in, nil
}

var msgChannel = make(chan pb.ChatMessage)

func (s *ChatServiceServer) Subscribe(in *pb.JoinMessage, srv pb.ChatService_SubscribeServer) error {
	fmt.Printf("Client connected: %v", in.Sender)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		for {
			defer wg.Done()
			fmt.Println("test!")
			message := <-msgChannel
			fmt.Printf("I GOT A MESSAGE IT WORKS!!")
			err := srv.Send(&message)
			if err != nil {
				log.Printf("send error %v", err)
			}
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &ChatServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
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
	mu       sync.Mutex
	channels []chan *pb.ChatMessage
}

func (s *ChatServiceServer) SendMessage(stream pb.ChatService_SendMessageServer) error {
	msg, err := stream.Recv()
	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}
	fmt.Println("I received a message")

	ack := pb.MessageAcknowledgement{
		T:               0,
		Acknowledgement: "Server received message",
	}
	stream.SendAndClose(&ack)

	fmt.Println(s.channels)
	for _, msgChan := range s.channels {
		msgChan <- msg
	}
	return nil
}

func (s *ChatServiceServer) Subscribe(joinMessage *pb.JoinMessage, stream pb.ChatService_SubscribeServer) error {
	fmt.Printf("Client connected: %v \n", joinMessage.Sender)

	msgChannel := make(chan *pb.ChatMessage)
	s.channels = append(s.channels, msgChannel)

	waitChannel := make(chan struct{})
	go func() {
		for {
			msg := <-msgChannel
			fmt.Printf("Received message: %v \n", msg)
			stream.Send(msg)
		}
	}()

	<-waitChannel
	return nil
}

func newServer() *ChatServiceServer {
	s := &ChatServiceServer{
		channels: make([]chan *pb.ChatMessage, 0),
	}
	return s
}

func main() {
	fmt.Println("--- SERVER APP ---")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, newServer())
	log.Printf("server listening at %v", lis.Addr())
	s.Serve(lis)

	/*for {
		 server :=<-subscribeChannel
		 go func(subscribeServer pb.ChatService_SubscribeServer) {
			 for {
				 fmt.Println("I AM RUNNING!!")
				 //defer wg.Done()
				 fmt.Println("test!")
				 message := <-msgChannel
				 fmt.Printf("I GOT A MESSAGE IT WORKS!!")
				 err := server.Send(&message)
				 if err != nil {
					 log.Printf("send error %v", err)
				 }
			 }
		 }(server)
	}*/
}

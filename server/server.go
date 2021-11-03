package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "mini_project_dis_sys"
	"net"
	"os"
	"strconv"
	"sync"
)

const (
	port = ":50051"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	mu      sync.Mutex
	clients []Client
}

type Client struct {
	senderName string
	channel    chan *pb.ChatMessage
	status     string
}

var T int64

func (s *ChatServiceServer) SendMessage(stream pb.ChatService_SendMessageServer) error {
	msg, err := stream.Recv()
	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}
	log.Println("I received a message")

	acknowledgementMessage := pb.MessageAcknowledgement{
		T:               T,
		Acknowledgement: "Sent message",
	}
	stream.SendAndClose(&acknowledgementMessage)

	// Update lamport timestamp
	if msg.T > T {
		T = msg.T
	}
	T++
	msg.T = T

	msg.Message = "(" + msg.Sender + "): " + msg.Message

	log.Printf("----UPDATED LAMPORT TIME: %v \n", T)
	go s.Broadcast(msg)
	return nil
}

func (s *ChatServiceServer) Broadcast(msg *pb.ChatMessage) {
	log.Println(s.clients)
	for _, client := range s.clients {
		if client.status != "disconnected" {
			log.Println("1111")
			log.Println(msg)
			log.Println(client.channel)
			client.channel <- msg
			log.Println("2222")
		}
	}
}

func (s *ChatServiceServer) Unsubscribe(ctx context.Context, leaveMessage *pb.LeaveMessage) (*pb.Null, error) {
	log.Printf("Client Disconnected: %v \n", leaveMessage.Sender)

	for i := 0; i < len(s.clients); i++ {
		if s.clients[i].senderName == leaveMessage.Sender {
			s.clients[i].status = "disconnected"
			T++
			log.Printf("----UPDATED LAMPORT TIME: %v \n", T)
			go s.Broadcast(&pb.ChatMessage{
				T:       T,
				Sender:  leaveMessage.Sender,
				Message: "Participant " + leaveMessage.Sender + " left Chitty-Chat at Lamport time " + strconv.FormatInt(T, 10),
			})
		}
	}

	return &pb.Null{}, nil
}

func (s *ChatServiceServer) Subscribe(joinMessage *pb.JoinMessage, stream pb.ChatService_SubscribeServer) error {
	log.Printf("Client connected: %v \n", joinMessage.Sender)

	client := Client{
		senderName: joinMessage.Sender,
		channel:    make(chan *pb.ChatMessage),
		status:     "connected",
	}

	s.clients = append(s.clients, client)

	if joinMessage.T > T {
		T = joinMessage.T
	}
	T++
	log.Printf("----UPDATED LAMPORT TIME: %v \n", T)

	go s.Broadcast(&pb.ChatMessage{
		T:       T,
		Sender:  joinMessage.Sender,
		Message: "Participant " + joinMessage.Sender + " joined Chitty-Chat at Lamport time " + strconv.FormatInt(T, 10),
	})

	waitChannel := make(chan struct{})

	go func() {
		for {
			msg := <-client.channel
			log.Printf("Received message: %v \n", msg)
			stream.Send(msg)
		}
	}()

	<-waitChannel
	return nil
}

func newServer() *ChatServiceServer {
	s := &ChatServiceServer{
		clients: make([]Client, 0),
	}
	return s
}

func main() {
	file, err := os.OpenFile("logFile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!")

	log.Println("--- SERVER APP ---")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("----UPDATED LAMPORT TIME: %v \n", T)

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, newServer())
	log.Printf("server listening at %v", lis.Addr())
	s.Serve(lis)
}

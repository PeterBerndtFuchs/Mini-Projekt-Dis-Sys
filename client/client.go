package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "mini_project_dis_sys"
	"os"
)

const (
	address = "localhost:50051"
)

var name string

func join(ctx context.Context, client pb.ChatServiceClient) {
	joinMessage := pb.JoinMessage{Sender: name}
	stream, err := client.Subscribe(ctx, &joinMessage)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Joined server\n")

	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				fmt.Println("I'm done")
				return
			}
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			fmt.Printf("(%v): %v \n", in.Sender, in.Message)
		}
	}()

	<-waitc
}

func sendMessage(ctx context.Context, client pb.ChatServiceClient, message string) {
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Can't send message: %v", err)
	}
	msg := pb.ChatMessage{
		T:       0,
		Message: message,
		Sender:  name,
	}
	stream.Send(&msg)

	_, err = stream.CloseAndRecv()
}

func leave(ctx context.Context, client pb.ChatServiceClient) {
	_, err := client.Unsubscribe(ctx, &pb.LeaveMessage{
		T:      0,
		Sender: name,
	})
	if err != nil {
		log.Printf("Can't send leave message: %v", err)
	}
}

func main() {

	// Connect
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Name
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your chat name:")
	scanner.Scan()
	name = scanner.Text()

	// Create stream
	client := pb.NewChatServiceClient(conn)
	go join(context.Background(), client)

	for scanner.Scan() {
		if scanner.Text() == "/leave" {
			go leave(context.Background(), client)
		} else {
			go sendMessage(context.Background(), client, scanner.Text())
		}
	}

}

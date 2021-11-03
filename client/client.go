package main

import (
	"bufio"
	"context"
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
var T int64

func join(ctx context.Context, client pb.ChatServiceClient) {
	T++
	joinMessage := pb.JoinMessage{
		T:      T,
		Sender: name,
	}
	log.Printf("----I JOINED WITH THE TIME: %v \n", T)

	stream, err := client.Subscribe(ctx, &joinMessage)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Joined server\n")

	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()

			if in.T > T {
				T = in.T
			}
			T++
			log.Printf("----I RECEIVED A MESSAGE WITH THE TIME: %v \n", T)

			if err == io.EOF {
				close(waitc)
				log.Println("I'm done")
				return
			}
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			log.Println(in.Message)
			// log.Printf("(%v): %v \n", in.Sender, in.Message)
		}
	}()

	<-waitc
}

func sendMessage(ctx context.Context, client pb.ChatServiceClient, message string) {
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Can't send message: %v", err)
	}

	T++
	msg := pb.ChatMessage{
		T:       T,
		Message: message,
		Sender:  name,
	}
	stream.Send(&msg)
	log.Printf("----I SENT A MESSAGE WITH THE TIME: %v \n", T)

	_, err = stream.CloseAndRecv()
}

func leave(ctx context.Context, client pb.ChatServiceClient) {
	T++
	_, err := client.Unsubscribe(ctx, &pb.LeaveMessage{
		T:      T,
		Sender: name,
	})
	log.Printf("----I LEFT WITH THE TIME: %v \n", T)
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
	log.Println("Enter your chat name:")
	scanner.Scan()
	name = scanner.Text()

	// Create stream
	client := pb.NewChatServiceClient(conn)
	go join(context.Background(), client)

	for scanner.Scan() {
		if scanner.Text() == "/leave" {
			go leave(context.Background(), client)
		} else {
			if len(scanner.Text()) <= 128 {
				go sendMessage(context.Background(), client, scanner.Text())
			} else {
				log.Println("Message is too long. Max length 128 characters")
			}
		}
	}

}

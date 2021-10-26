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
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your chat name:")
	scanner.Scan()
	name := scanner.Text()

	// create stream
	client := pb.NewChatServiceClient(conn)
	in := &pb.JoinMessage{Sender: name}
	stream, err := client.Subscribe(context.Background(), in)
	if err != nil {
		log.Fatalf("Subscribe error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Message)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")

	for {
		scanner.Scan()

		if len(scanner.Text()) <= 128 {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			_, err := c.Publish(ctx, &pb.ChatMessage{
				T:       0,
				Sender:  name,
				Message: scanner.Text(),
			})
			if err != nil {
				log.Printf("could not send the message: %v", err)
			}
		} else {
			fmt.Println("Max message length is 128 characters")
		}
	}

}

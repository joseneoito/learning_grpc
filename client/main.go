package main

import (
	pb "blogs/client/pb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogsClient(conn)
	blogList, err := client.GetBlogList(context.Background(), &pb.GetBlogListRequest{})
	if err != nil {
		log.Fatalf("failed to get blog list: %v", err)
	}
	log.Printf("blog list: %v", blogList)
}

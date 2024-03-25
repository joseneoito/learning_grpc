package main

import (
    "net"
    "log"
    "context"
    pb "blogs/server/pb"
    "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

)

type server struct {
    pb.UnimplementedBlogsServer
}

func (s *server)GetBlogList(ctx context.Context, in *pb.GetBlogListRequest)(*pb.GetBlogListResponse, error){
    //log.Printf("received request %v", in.ProtoReflect().Descriptor().FullName())
    return &pb.GetBlogListResponse{
        Blogs: getSampleBlogs(),
    }, nil
}

func main(){
    listen, err := net.Listen("tcp",":8080"); 
    if err != nil{
        panic(err)
    }
    s := grpc.NewServer();
    reflection.Register(s);
    pb.RegisterBlogsServer(s, &server{})
    if err := s.Serve(listen); err !=nil{
        log.Fatalf("Faield to server: %v", err)
    }


}
func getSampleBlogs() []*pb.Blog {
    blogs := []*pb.Blog{
        {Slug: "slug1", Title: "title1", Author: "author1", Keywords: 42},
        {Slug: "slug2", Title: "title2", Author: "author2", Keywords: 24},
        {Slug: "slug3", Title: "title3", Author: "author3", Keywords: 36},
        {Slug: "slug4", Title: "title4", Author: "author4", Keywords: 48},
    }
    return blogs

}

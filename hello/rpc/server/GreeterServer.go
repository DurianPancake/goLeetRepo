package server

import (
	"context"
	pb "goFirst/role/api/rpc/proto"
	//pb2 "goSecond/hello/rpc/proto"
)

type server struct {

}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest)(*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
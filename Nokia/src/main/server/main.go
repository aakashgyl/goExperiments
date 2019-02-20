package main

import (
	"log"
	"net"
	"io"
	"fmt"
	"google.golang.org/grpc"
	pb "main/proto"
	"google.golang.org/grpc/reflection"
	"os"
)

const (
	port = ":50051"
)

// server is used to implement 
type server struct{
	cmdata []*pb.CmData
}

// SayHello implements ndac.EdgeServices
func (s *server) SendCMData(stream pb.EdgeServices_SendCMDataServer) error {
	var cmdata []byte
	f, err := os.Create("C:/Users/aagoyal/eclipse-workspace-oxygen/NDACServer/log.txt")
    check(err)
    defer f.Close()
    
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			_, err = f.Write(cmdata[:])
		    check(err)
			return stream.SendAndClose(&pb.CmReply{
				Status: "SUCCESS",
				Message: "Working fine",
			})
		}
		
		if err != nil {
			return err
		}
		cmdata = append(cmdata, point.GetData()...)
		fmt.Println(string(point.GetData()))
	}	
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	
	pb.RegisterEdgeServicesServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
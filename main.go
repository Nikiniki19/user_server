package main

import (
	"log"
	"net"
	"sync"
	"userservice/handler"
	"userservice/proto"

	"google.golang.org/grpc"
)

const (
	port1 = ":8083"
)
const (
	port2 = ":8084"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go grpcConnectionClient1(port1, wg)
	go grpcConnectionClient2(port2, wg)
	wg.Wait()
}
func grpcConnectionClient1(port1 string, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", port1)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterClient1RequestServer(grpcServer, &handler.CreateUser{})
	log.Printf("Server started at : %v", lis.Addr())

	err1 := grpcServer.Serve(lis)
	if err1 != nil {
		log.Fatalf("Failed to start: %v", err1)
	}
}

func grpcConnectionClient2(port2 string, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", port2)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterClient2RequestServer(grpcServer, &handler.FetchUser{})
	log.Printf("Server started at : %v", lis.Addr())

	err1 := grpcServer.Serve(lis)
	if err1 != nil {
		log.Fatalf("Failed to start: %v", err1)
	}

}

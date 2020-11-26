package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RaminCH/go3_grpc/Lec3/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Starting gRPC server...")

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen from port: %v", err)
	}

	//Our Service
	serv := chat.Server{} // nash servis kotoriy umeyet delat te deystviya kotoriye propisani v kontrakte(chat.proto-> service ChatService {...)
	//pochemu mi mojem utverjdat chto on mojet vipolnat deystviya SayHello i SaygoodBye v kontrakte, potomu chto mi v ruchnuyu
	//v chat.go propisali eti metodi(Sayhe.. Saygood..), chtobi serv udovletvoral interfeysu -> type ChatServiceServer interface
	// v chat.pb.go

	//basic gRPC server
	grpcServer := grpc.NewServer() //grpc server (grubo govora - eto router kak v REST)

	//Assotsiiiruyem gRPC server s nashim servisom(serv)
	chat.RegisterChatServiceServer(grpcServer, &serv)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve with that point: %s", err)
	}

}

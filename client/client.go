package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RaminCH_self/Go3_gRPC/lec3/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client starting ...")

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8000", grpc.WithInsecure()) //soyedinayemsa k portu servera (on toje 8000)
	if err != nil {
		log.Fatalf("cannot connect to: %v", err)
	}

	defer conn.Close() //client mojet podkluchatsa i otkluchatsa

	cli := chat.NewChatServiceClient(conn) //sozdayem obyekt-(kliyent) kotorogo soyedinayem k portu 8000
	//NewChatServiceClient iz chat.pb.go gde on realizovivayet takje Sayhello i SayGoodbye

	//Dergayem Sayhello s servera
	resp, err := cli.SayHello(context.Background(), &chat.Message{Body: "Hello from Client!"}) //message from client to server
	if err != nil {
		log.Fatalf("Error when SayHello: %s", err)
	}
	log.Printf("Response from Server: %s", resp.Body) //response from server

	//Dergayem SayGoodBye u servera ot client-a
	resp, err = cli.SayGoodBye(context.Background(), &chat.Message{Body: "GoodBye from Client!"}) //message from client to server
	if err != nil {
		log.Fatalf("Error when SayGoodBye: %s", err)
	}
	log.Printf("Response from Server: %s", resp.Body) //response from server
}

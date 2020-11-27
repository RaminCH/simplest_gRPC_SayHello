package chat

import (
	context "context"
	"log"
)

//Server
type Server struct { // udovletvorayet interfeysu ->>>> type ChatServiceServer interface {..}, potomu chto tam ne realizovano nichego
	//...
}

//importing both SayHello and SayGoodBye from chat.pb.go -> type ChatServiceServer interface
//SayHello...
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", in.Body) //Body is in the type Message struct {... in chat.pb.proto
	return &Message{Body: "Hello from this server!"}, nil        // Message struct is in the chat.pb.proto -> type Message struct {..
}

//SayGoodBye...
func (s *Server) SayGoodBye(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", in.Body)
	return &Message{Body: "GoodBye from this server"}, nil
}

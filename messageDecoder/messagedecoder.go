package messagedecoder

import (
	"PolitoGroupHelpBot/utils"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type MessageDecoder struct {
	UnimplementedMessageDecoderServer
}

func (s *MessageDecoder) Decode(ctx context.Context, message *TgMessageInfo) (*DecoderResult, error) {
	fmt.Println(message.MessageInfo)
	return &DecoderResult{Ok: true}, nil
}

func Main() {
	lis, err := net.Listen("tcp", utils.LoadPortFromEnv("DECODER_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dec := MessageDecoder{}

	server := grpc.NewServer()
	RegisterMessageDecoderServer(server, &dec)

	log.Printf("Server started on port %v\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Could not start RPC server: %v", err)
	}
}

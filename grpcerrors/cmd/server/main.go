package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ellemouton/playgrounds/grpcerrors"
	"github.com/ellemouton/playgrounds/grpcerrors/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = flag.Int64("port", 8080, "the port that the server will listen on")

func main() {
	flag.Parse()

	addr := fmt.Sprintf("localhost:%d", *port)
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grcpServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	server := &grpcerrors.Server{}

	pb.RegisterErrorsServer(grcpServer, server)

	log.Printf("Server listening on %s\n", addr)
	log.Fatalln(grcpServer.Serve(list))
}

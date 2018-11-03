package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/kjbreil/keyboard/keyrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 10000, "The server port")
)

type server struct{}

func newServer() *server {
	s := &server{}
	return s
}

func (s *server) KeyRoute(stream pb.KeyRPC_KeyRouteServer) error {

	fmt.Println("hjerer")

	for {
		key, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("I would be pressing this Name: %s Virtual: %d Scan: %d", key.KeyName, key.Virtual, key.Scan)
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = "server1.pem"
		}
		if *keyFile == "" {
			*keyFile = "server1.key"
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterKeyRPCServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

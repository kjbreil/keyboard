package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/kjbreil/keyboard"

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
	// startTime := time.Now()

	for {
		key, err := stream.Recv()
		if err == io.EOF {
			// endTime := time.Now()
			return stream.SendAndClose(&pb.EntrySummary{
				Error: "",
			})
		}
		if err != nil {
			return stream.SendAndClose(&pb.EntrySummary{
				Error: err.Error(),
			})
		}
		log.Printf("I am pressing this: %s Virtual: %d Scan: %d", key.KeyName, key.Virtual, key.Scan)

		kp := keyboard.KeyPress{
			Key: rune(key.Virtual),
		}
		err = kp.Press()
		if err != nil {
			log.Fatalf("There was an errur: %v\n", err)
		}
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

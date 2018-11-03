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

func (s *server) KeyBurst(stream pb.KeyRPC_KeyBurstServer) error {
	// startTime := time.Now()

	for {
		key, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF MET")
			// endTime := time.Now()
			return stream.SendAndClose(&pb.Summary{
				Complete: true,
			})
		}
		if err != nil {
			log.Println(err)
			return stream.SendAndClose(&pb.Summary{
				Complete: false,
			})
		}

		kp := keyboard.KeyPress{
			Key:   rune(key.Key),
			Upper: key.Upper,
		}
		if int(key.Sleep) != 0 {
			sl := int(key.Sleep)
			kp.Sleep = &sl
		}
		if int(key.Modifier) != 0 {
			mo := rune(key.Modifier)
			kp.Modifier = &mo
		}
		log.Printf("Going to press: %s\n", keyboard.Scan[kp.Key].Name)
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

	// opts = append(opts, grpc.ConnectionTimeout(600*time.Second))
	log.Println(opts)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterKeyRPCServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

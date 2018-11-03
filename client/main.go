package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/kjbreil/keyboard"
	pb "github.com/kjbreil/keyboard/keyrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func runKey(client pb.KeyRPCClient) {
	var keys []*pb.Key
	for i := 0; i < 10; i++ {
		keys = append(keys, randomKey())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.KeyRoute(ctx)
	if err != nil {
		log.Fatalf("%v.KeyRoute(_) = _, %v", client, err)
	}
	for _, key := range keys {
		log.Printf("Sending Key: %s", key.KeyName)
		if err := stream.Send(key); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, key, err)
		}
	}
	reply, err := stream.CloseAndRecv()

	if reply.Error != "" {
		log.Fatalf("Got Error from Server: %s", reply.Error)
	}
}

func randomKey() *pb.Key {
	key := keyboard.Scan[randKey(keyboard.Scan)]

	return &pb.Key{
		KeyName: key.Name,
		Virtual: uint32(key.Virtual),
		Scan:    uint32(key.Scan),
		Sleep:   100,
		Mock:    true,
	}

}

func randKey(scan map[rune]keyboard.VirtScan) rune {
	for k := range scan {
		return k
	}
	return 0
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = "ca.pem"
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewKeyRPCClient(conn)

	runKey(client)
}

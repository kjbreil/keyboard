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
	keys := stringToKeys("111112222233333444445555566666777778888899999")
	ctx, cancel := context.WithTimeout(context.Background(), 360*time.Second)
	defer cancel()
	stream, err := client.KeyRoute(ctx)
	if err != nil {
		log.Fatalf("%v.KeyRoute(_) = _, %v", client, err)
	}
	for _, key := range keys {
		time.Sleep(time.Duration(key.Sleep) * time.Millisecond)
		log.Printf("Sending Key: %s", key.KeyName)
		err := stream.Send(key)
		if err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, key, err)
		}
	}
	reply, err := stream.CloseAndRecv()

	if !reply.Complete {
		log.Fatalf("Got Error from Server:")
	}
}

func randomKey() *pb.Key {
	key := keyboard.Scan[randKey(keyboard.Scan)]

	return &pb.Key{
		KeyName: key.Name,
		Virtual: uint32(key.Virtual),
		Scan:    uint32(key.Scan),
		Sleep:   10,
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

func stringToKeys(s string) (keys []*pb.Key) {
	ks, bs := 100, 100
	// silently throwing away an error - get rid of this ASAP
	b, _ := keyboard.StringToBurst(s, &ks, &bs)
	for _, k := range b.Presses {
		var key = &pb.Key{
			KeyName: keyboard.Scan[k.Key].Name,
			Virtual: uint32(keyboard.Scan[k.Key].Virtual),
			Scan:    uint32(keyboard.Scan[k.Key].Scan),
			Sleep:   500,
			Mock:    true,
		}
		keys = append(keys, key)
		// k.
	}
	return
}

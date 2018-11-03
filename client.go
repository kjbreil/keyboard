package keyboard

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/kjbreil/keyboard/keyrpc"
	"google.golang.org/grpc"
)

// RunBurst sends a complete burst
func RunBurst(b KeyBurst, serverAddr *string) error {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewKeyRPCClient(conn)

	keys := burstToKeys(b)
	ctx, cancel := context.WithTimeout(context.Background(), 360*time.Second)
	defer cancel()
	stream, err := client.KeyBurst(ctx)
	if err != nil {
		return err
	}
	for _, key := range keys {
		time.Sleep(time.Duration(key.Sleep) * time.Millisecond)
		log.Printf("Sending Key: %s", Scan[rune(key.Key)].Name)
		err := stream.Send(key)
		if err != nil {
			return err
		}
	}
	reply, err := stream.CloseAndRecv()

	if !reply.Complete {
		return fmt.Errorf("got Error from Server")
	}
	return nil
}

func randomKey() *pb.KeyPress {
	key := Scan[randKey(Scan)]

	return &pb.KeyPress{
		Key:   uint32(key.Virtual),
		Sleep: 10,
	}

}

func randKey(scan map[rune]VirtScan) rune {
	for k := range scan {
		return k
	}
	return 0
}

func burstToKeys(b KeyBurst) (keys []*pb.KeyPress) {
	for _, k := range b.Presses {
		var key = &pb.KeyPress{
			Key:   uint32(k.Key),
			Upper: k.Upper,
		}
		if k.Modifier != nil {
			key.Modifier = uint32(*k.Modifier)
		}
		if k.Sleep != nil {
			key.Sleep = uint32(*k.Sleep)
		}
		keys = append(keys, key)
		// k.
	}
	return
}

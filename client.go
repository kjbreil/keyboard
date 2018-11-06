package keyboard

import (
	"context"
	"fmt"
	"time"

	pb "github.com/kjbreil/keyboard/keyrpc"
	"google.golang.org/grpc"
)

// Server sends a complete sequence to a server
func (ks KeySeq) Server(serverAddr *string) error {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewKeyRPCClient(conn)

	var keys []*pb.KeyPress
	for _, eb := range ks.Bursts {
		keys = append(keys, burstToKeys(eb)...)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.KeyBurst(ctx)
	if err != nil {
		return err
	}
	for _, key := range keys {
		err := stream.Send(key)
		if err != nil {
			return err
		}
	}
	reply, err := stream.CloseAndRecv()

	if reply != nil && !reply.Complete {
		return fmt.Errorf("got Error from Server")
	}
	if err != nil {
		return err
	}
	return nil
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
	}
	return
}

// Server sends a complete burst to a server
func (b KeyBurst) Server(serverAddr *string) error {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewKeyRPCClient(conn)

	keys := burstToKeys(b)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.KeyBurst(ctx)
	if err != nil {
		return err
	}
	for _, key := range keys {
		err := stream.Send(key)
		if err != nil {
			return err
		}
	}
	reply, err := stream.CloseAndRecv()

	if reply != nil && !reply.Complete {
		return fmt.Errorf("got Error from Server")
	}
	if err != nil {
		return err
	}
	return nil
}

// ServerSwitchWindow switches to a windows with name windowName, must be exact
func ServerSwitchWindow(windowName string, serverAddr *string) error {
	wn := new(pb.WindowName)
	wn.Name = windowName
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewKeyRPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sum, err := client.SwitchWindow(ctx, wn)
	if sum != nil && sum.Complete {
		return nil
	} else if err != nil {
		return err
	} else {
		return fmt.Errorf("Window switch error")
	}

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

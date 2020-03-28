package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	pb "grpc-lb/addr"
	"os"
	"time"
)

const (
	TIMEOUT       = 5 * time.Second
	WAIT_DURATION = 1 * time.Second
)

var counter int

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <host:port>\n", os.Args[0])
	os.Exit(2)
}

func callAndShowResponse(client pb.AddrClient) {
	reply, err := client.GetAddr(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	counter++
	fmt.Printf("#%d: %s\n", counter, reply.Addr)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	hostPort := os.Args[1]

	fmt.Printf("Connecting to %s...\n", hostPort)
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Printf("Making rpc...\n")
	client := pb.NewAddrClient(conn)
	for {
		callAndShowResponse(client)
		time.Sleep(WAIT_DURATION)
	}
}

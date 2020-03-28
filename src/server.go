package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-lb/addr"
	"net"
	"net/http"
)

const (
	PLAIN_HTTP_PORT = ":80"
	GRPC_PORT       = ":30051"
)

var localIp string

type serverGrpc struct{}

func (s *serverGrpc) GetAddr(ctx context.Context, in *empty.Empty) (*pb.AddrReply, error) {
	return &pb.AddrReply{Addr: localIp}, nil
}

// getOutboundIP gets preferred outbound ip of this machine.
// @see https://stackoverflow.com/a/37382208/714426
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func addr(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, localIp)
}

func health(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "OK")
}

func main() {
	localIp = getOutboundIP().String()

	go func() {
		log.Infof("Starting http server on %s...", PLAIN_HTTP_PORT)
		http.HandleFunc("/addr", addr)
		http.HandleFunc("/health", health)
		if err := http.ListenAndServe(PLAIN_HTTP_PORT, nil); err != nil {
			log.Fatalf("Fail to start HTTP server: %v", err)
		}
	}()

	log.Infof("Starting gRPC server on %s...", GRPC_PORT)
	lis, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		log.Fatalf("Fail to start gRPC server: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddrServer(s, &serverGrpc{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to start gRPC server: %v", err)
	}
}

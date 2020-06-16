package main

import (
	protos "github.com/go-grpc-basics/protos/currency"
	"github.com/go-grpc-basics/protos/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	_ = gs.Serve(l)
}
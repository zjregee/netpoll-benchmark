package main

import (
	"context"

	"github.com/zjregee/anet"

	"github.com/cloudwego/netpoll-benchmark/anet/codec"
	"github.com/cloudwego/netpoll-benchmark/runner"
)

func NewRPCServer() runner.Server {
	return &rpcServer{}
}

var _ runner.Server = &rpcServer{}

type rpcServer struct{}

func (s *rpcServer) Run(network, address string) error {
	listener, err := anet.CreateListener(network, address)
	if err != nil {
		panic(err)
	}
	eventLoop, err := anet.NewEventLoop(s.handler)
	if err != nil {
		panic(err)
	}
	return eventLoop.Serve(listener)
}

func (s *rpcServer) handler(_ context.Context, conn anet.Connection) (err error) {
	reader, writer := conn.Reader(), conn.Writer()
	req := &runner.Message{}
	err = codec.Decode(reader, req)
	if err != nil {
		return err
	}
	resp := runner.ProcessRequest(reporter, req)
	err = codec.Encode(writer, resp)
	if err != nil {
		return err
	}
	return nil
}

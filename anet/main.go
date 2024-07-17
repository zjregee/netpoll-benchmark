package main

import (
	"github.com/cloudwego/netpoll-benchmark/runner"
	"github.com/cloudwego/netpoll-benchmark/runner/perf"
	"github.com/cloudwego/netpoll-benchmark/runner/svr"
)

func main() {
	svr.Serve(NewServer)
}

var reporter = perf.NewRecorder("ANET@Server")

func NewServer(mode runner.Mode) runner.Server {
	switch mode {
	case runner.Mode_Echo:
		return NewRPCServer()
	default:
		return nil
	}
}

package main

import (
	"github.com/saushew/huobiRAW/cmd/accountclientexample"
	"github.com/saushew/huobiRAW/cmd/accountwebsocketclientexample"
	"github.com/saushew/huobiRAW/cmd/algoorderclientexample"
	"github.com/saushew/huobiRAW/cmd/commonclientexample"
	"github.com/saushew/huobiRAW/cmd/crossmarginclientexample"
	"github.com/saushew/huobiRAW/cmd/etfclientexample"
	"github.com/saushew/huobiRAW/cmd/isolatedmarginclientexample"
	"github.com/saushew/huobiRAW/cmd/marketclientexample"
	"github.com/saushew/huobiRAW/cmd/marketwebsocketclientexample"
	"github.com/saushew/huobiRAW/cmd/orderclientexample"
	"github.com/saushew/huobiRAW/cmd/orderwebsocketclientexample"
	"github.com/saushew/huobiRAW/cmd/stablecoinclientexample"
	"github.com/saushew/huobiRAW/cmd/subuserclientexample"
	"github.com/saushew/huobiRAW/cmd/walletclientexample"
	"github.com/saushew/huobiRAW/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}

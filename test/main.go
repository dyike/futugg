package main

import (
	"fmt"
	"futugg"
	_ "futugg/handlers"
	// "futugg/pb/InitConnect"
	// "github.com/golang/protobuf/proto"
)

func main() {

	block := make(chan bool)
	cli := futugg.New("0.0.0.0", "11250", "")

	futugg.Cmd("send.GetGlobalState", cli)
	cli.KeepAlive()

	// recv
	go func() {
		fmt.Println("start recv data")
		cli.Recv()
	}()

	futugg.Cmd("send.Qot_Sub", cli, "HK.01810", "Basic", true, true, false)
	// futugg.Cmd("send.Qot_RegQotPush", cli, "US.BILI", "Basic", true, false)
	// futugg.Cmd("send.Qot_GetSubInfo", cli, true)
	futugg.Cmd("recv.Qot_UpdateBasicQot", cli, "HK.01810")

	<-block
}

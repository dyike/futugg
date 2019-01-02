package handlers

import (
	"fmt"
	"futugg"
	"futugg/pb/Notify"
	"log"

	"github.com/golang/protobuf/proto"
)

func init() {
	futugg.SetHandlerId(uint32(1003), "Notify")

	var err error
	err = futugg.On("recv.Notify", NotifyRecv)
	if err != nil {
		log.Fatalln(err)
	}
}

func NotifyRecv(data []byte) error {
	resp := &Notify.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return fmt.Error("marshal error: %s", err)
	}

	return nil
}

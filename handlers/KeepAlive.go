package handlers

import (
	"fmt"
	"futugg"
	"futugg/pb/KeepAlive"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

func init() {
	futugg.SetHandlerId(uint32(1004), "KeepAlive")

	var err error
	err = futugg.On("send.KeepAlive", KeepAliveSend)
	err = futugg.On("recv.KeepAlive", KeepAliveRecv)
	if err != nil {
		log.Fatalln(err)
	}
}

func KeepAliveSend(conn *futugg.FutuGG) error {
	pack := &futugg.FutuPack{}
	pack.SetProto(uint32(1004))

	time := time.Now().Unix()

	reqData := &KeepAlive.Request{
		C2S: &KeepAlive.C2S{
			Time: &time,
		},
	}

	pbData, err := proto.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	pack.SetBody(pbData)
	err = conn.Send(pack)

	return err
}

func KeepAliveRecv(data []byte) error {
	resp := &KeepAlive.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	m := jsonpb.Marshaler{}
	result, err := m.MarshalToString(resp)
	fmt.Println(result)
    return err
}

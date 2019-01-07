package handlers

import (
	"fmt"
	"futugg"
	"futugg/pb/Qot_GetSubInfo"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

func init() {
	futugg.SetHandlerId(uint32(3003), "Qot_GetSubInfo")
	var err error
	err = futugg.On("send.Qot_GetSubInfo", QotGetSubInfoSend)
	err = futugg.On("recv.Qot_GetSubInfo", QotGetSubInfoRecv)
	if err != nil {
		fmt.Println(err)
	}
}

func QotGetSubInfoSend(conn *futugg.FutuGG, isReqAllConn bool) error {
	pack := &futugg.FutuPack{}
	pack.SetProto(uint32(3003))

	reqData := &Qot_GetSubInfo.Request{
		C2S: &Qot_GetSubInfo.C2S{
			IsReqAllConn: &isReqAllConn,
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

func QotGetSubInfoRecv(data []byte) error {
	resp := &Qot_GetSubInfo.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	m := jsonpb.Marshaler{}
	result, err := m.MarshalToString(resp)
	fmt.Println(result)

	return err
}

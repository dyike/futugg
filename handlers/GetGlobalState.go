package handlers

import (
	"fmt"
	"futugg"
	"futugg/pb/GetGlobalState"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func init() {
	futugg.SetHandlerId(uint32(1002), "GetGlobalState")
	var err error
	err = futugg.On("send.GetGlobalState", GetGlobalStateSend)
	err = futugg.On("recv.GetGlobalState", GetGlobalStateRecv)
	if err != nil {
		fmt.Println(err)
	}
}

// GetGlobalStateSend handler
func GetGlobalStateSend(conn *futugg.FutuGG) error {

	pack := &futugg.FutuPack{}

	pack.SetProto(uint32(1002))
	userID := conn.LoginUserID

	reqData := &GetGlobalState.Request{
		C2S: &GetGlobalState.C2S{
			UserID: &userID,
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

// GetGlobalStateRecv handler
func GetGlobalStateRecv(data []byte) (string, error) {
	resp := &GetGlobalState.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return "", fmt.Errorf("marshal error: %s", err)
	}

	m := jsonpb.Marshaler{}
	result, err := m.MarshalToString(resp)
	return result, nil
}

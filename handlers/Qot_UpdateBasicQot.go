package handlers

import (
	// "encoding/json"
	"fmt"
	"futugg"
	"futugg/pb/Qot_UpdateBasicQot"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

func init() {
	futugg.SetHandlerId(uint32(3005), "Qot_UpdateBasicQot")

	var err error
	err = futugg.On("recv.Qot_UpdateBasicQot", QotUpdateBasicQotRecv)
	if err != nil {
		fmt.Println(err)
	}
}

func QotUpdateBasicQotRecv(data []byte) (string, error) {
	resp := &Qot_UpdateBasicQot.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

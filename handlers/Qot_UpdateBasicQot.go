package handlers

import (
	// "encoding/json"
	"fmt"
	"futugg"
	"futugg/pb/Qot_UpdateBasicQot"

	"github.com/golang/protobuf/proto"
)

func init() {
	futugg.SetHandlerId(uint32(3005), "Qot_UpdateBasicQot")

	var err error
	err = futugg.On("recv.Qot_UpdateBasicQot", QotUpdateBasicQotRecv)
	if err != nil {
		fmt.Println(err)
	}
}

func QotUpdateBasicQotRecv(data []byte) error {
	resp := &Qot_UpdateBasicQot.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	return nil
}
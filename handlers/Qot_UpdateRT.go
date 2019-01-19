package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_UpdateRT"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3009), "Qot_UpdateRT")

    var err error
    err = futugg.On("recv.Qot_UpdateRT", QotUpdateRTRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotUpdateRTRecv(data []byte) (string, error) {
    resp := &Qot_UpdateRT.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}
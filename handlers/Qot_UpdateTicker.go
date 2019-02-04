package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_UpdateTicker"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3011), "Qot_UpdateTicker")

    var err error
    err = futugg.On("recv.Qot_UpdateTicker", QotUpdateTickerRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotUpdateTickerRecv(data []byte) error {
    resp := &Qot_UpdateTicker.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}
package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_UpdateTicker"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(3009), "Qot_UpdateTicker")

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

    return nil
}
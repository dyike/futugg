package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_UpdateKL"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(3007), "Qot_UpdateKL")

    var err error
    err = futugg.On("recv.Qot_UpdateKL", QotUpdateKLRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotUpdateKLRecv(data []byte) error {
    resp := &Qot_UpdateKL.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    return nil
}
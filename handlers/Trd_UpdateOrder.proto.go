package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_UpdateOrder"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(2208), "Trd_UpdateOrder")
    var err error
    err = futugg.On("recv.Trd_UpdateOrder", TrdUpdateOrderRecv)
    if err != nil {
        fmt.Println(err)
    }
}


func TrdUpdateOrderRecv(data []byte) error {
    resp := &Trd_UpdateOrder.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    fmt.Println(resp)

    return nil
}

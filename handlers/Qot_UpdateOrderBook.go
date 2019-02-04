package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_UpdateOrderBook"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3013), "Qot_UpdateOrderBook")

    var err error
    err = futugg.On("recv.Qot_UpdateOrderBook", QotUpdateOrderBookRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotUpdateOrderBookRecv(data []byte) error {
    resp := &Qot_UpdateOrderBook.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}
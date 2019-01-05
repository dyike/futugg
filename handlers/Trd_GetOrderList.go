package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_GetOrderList"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(2201), "Trd_GetOrderList")
    var err error
    err = futugg.On("send.Trd_GetOrderList", TrdGetOrderListSend)
    err = futugg.On("recv.Trd_GetOrderList", TrdGetOrderListRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func TrdGetOrderListSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32, filterStatusList []int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2201))

    header := setTrdHeader(trdEnv, accID, trdMarket)
    reqData := &Trd_GetOrderList.Request{
        C2S: &Trd_GetOrderList.C2S{
            Header: header,
            FilterStatusList: filterStatusList,
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

func TrdGetOrderListRecv(data []byte) error {
    resp := &Trd_GetOrderList.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    fmt.Println(resp)

    return nil
}

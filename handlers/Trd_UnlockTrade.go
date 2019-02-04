package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_UnlockTrade"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2005), "Trd_UnlockTrade")
    var err error
    err = futugg.On("send.Trd_UnlockTrade", TrdUnlockTradeSend)
    err = futugg.On("recv.Trd_UnlockTrade", TrdUnlockTradeRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func TrdUnlockTradeSend(conn *futugg.FutuGG, unlock bool) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2005))

    reqData := &Trd_UnlockTrade.Request{
        C2S: &Trd_UnlockTrade.C2S{
            Unlock: &unlock,
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

func TrdUnlockTradeRecv(data []byte) error {
    resp := &Trd_UnlockTrade.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}

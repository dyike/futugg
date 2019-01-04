package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetRT"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(3004), "Qot_GetRT")
    var err error
    err = futugg.On("send.Qot_GetRT", QotGetRTSend)
    err = futugg.On("recv.Qot_GetRT", QotGetRTRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetRTSend(conn *futugg.FutuGG, stockCode string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3008))

    security := transStockCode(stockCode)

    reqData := &Qot_GetRT.Request{
        C2S: &Qot_GetRT.C2S{
            Security: security,
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

func QotGetRTRecv(data []byte) error {
    resp := &Qot_GetRT.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    fmt.Println(resp)

    return nil
}

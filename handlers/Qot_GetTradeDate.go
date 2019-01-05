package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetTradeDate"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(3200), "Qot_GetTradeDate")
    var err error
    err = futugg.On("send.Qot_GetTradeDate", QotGetTradeDateSend)
    err = futugg.On("recv.Qot_GetTradeDate", QotGetTradeDateRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// market - HK: 1, US: 11, SH: 21, SZ: 22
func QotGetTradeDateSend(conn *futugg.FutuGG, marketNum int32, beginTime string, endTime string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3200))

    reqData := &Qot_GetTradeDate.Request{
        C2S: &Qot_GetTradeDate.C2S{
            Market: &marketNum,
            BeginTime: &beginTime,
            EndTime: &endTime,
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

func QotGetTradeDateRecv(data []byte) error {
    resp := &Qot_GetTradeDate.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    fmt.Println(resp)

    return nil
}

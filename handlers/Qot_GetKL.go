package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetKL"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(3006), "Qot_GetKL")

    var err error
    err = futugg.On("send.Qot_GetKL", QotGetKLSend)
    err = futugg.On("recv.Qot_GetKL", QotGetKLRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO KLine setting
func QotGetKLSend(conn *futugg.FutuGG, rehabType int32, klType int32, stockCode string, reqNum int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3006))

    security := transStockCode(stockCode)

    reqData := &Qot_GetKL.Request{
        C2S: &Qot_GetKL.C2S{
            RehabType: &rehabType,
            KlType: &klType,
            Security: security,
            ReqNum: &reqNum,
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

func QotGetKLRecv(data []byte) error {
    resp := &Qot_GetKL.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    return nil
}


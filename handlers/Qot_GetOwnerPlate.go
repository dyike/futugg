package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_Common"
    "futugg/pb/Qot_GetOwnerPlate"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3207), "Qot_GetOwnerPlate")
    var err error
    err = futugg.On("send.Qot_GetOwnerPlate", QotGetOwnerPlateSend)
    err = futugg.On("recv.Qot_GetOwnerPlate", QotGetOwnerPlateRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetOwnerPlateSend(conn *futugg.FutuGG, stockCode string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3207))

    var securityList []*Qot_Common.Security
    security := transStockCode(stockCode)
    securityList = append(securityList, security)

    reqData := &Qot_GetOwnerPlate.Request{
        C2S: &Qot_GetOwnerPlate.C2S{
            SecurityList:         securityList,
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

func QotGetOwnerPlateRecv(data []byte) error {
    resp := &Qot_GetOwnerPlate.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}

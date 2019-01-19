package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetPlateSecurity"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3205), "Qot_GetPlateSecurity")
    var err error
    err = futugg.On("send.Qot_GetPlateSecurity", QotGetPlateSecuritySend)
    err = futugg.On("recv.Qot_GetPlateSecurity", QotGetPlateSecurityRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetPlateSecuritySend(conn *futugg.FutuGG, stockCode string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3205))

    security := transStockCode(stockCode)

    reqData := &Qot_GetPlateSecurity.Request{
        C2S: &Qot_GetPlateSecurity.C2S{
            Plate: security,
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

func QotGetPlateSecurityRecv(data []byte) (string, error) {
    resp := &Qot_GetPlateSecurity.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

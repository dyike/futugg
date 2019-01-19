package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetReference"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3206), "Qot_GetReference")
    var err error
    err = futugg.On("send.Qot_GetReference", QotGetReferenceSend)
    err = futugg.On("recv.Qot_GetReference", QotGetReferenceRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetReferenceSend(conn *futugg.FutuGG, stockCode string, referenceType int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3206))

    security := transStockCode(stockCode)

    reqData := &Qot_GetReference.Request{
        C2S: &Qot_GetReference.C2S{
            Security: security,
            ReferenceType: &referenceType,
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

func QotGetReferenceRecv(data []byte) (string, error) {
    resp := &Qot_GetReference.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

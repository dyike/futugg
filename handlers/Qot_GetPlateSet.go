package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetPlateSet"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3204), "Qot_GetPlateSet")
    var err error
    err = futugg.On("send.Qot_GetPlateSet", QotGetPlateSetSend)
    err = futugg.On("recv.Qot_GetPlateSet", QotGetPlateSetRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetPlateSetSend(conn *futugg.FutuGG, marketNum int32, plateSetType int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3204))

    reqData := &Qot_GetPlateSet.Request{
        C2S: &Qot_GetPlateSet.C2S{
            Market: &marketNum,
            PlateSetType: &plateSetType,
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

func QotGetPlateSetRecv(data []byte) (string, error) {
    resp := &Qot_GetPlateSet.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetHoldingChangeList"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3208), "Qot_GetHoldingChangeList")
    var err error
    err = futugg.On("send.Qot_GetHoldingChangeList", QotGetHoldingChangeListSend)
    err = futugg.On("recv.Qot_GetHoldingChangeList", QotGetHoldingChangeListRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetHoldingChangeListSend(conn *futugg.FutuGG, stockCode string, holderCategory  int32, beginTime string, endTime string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3208))

    security := transStockCode(stockCode)
    reqData := &Qot_GetHoldingChangeList.Request{
        C2S: &Qot_GetHoldingChangeList.C2S{
            Security:         security,
            HolderCategory:   &holderCategory,
            BeginTime:        &beginTime,
            EndTime:          &endTime,
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

func QotGetHoldingChangeListRecv(data []byte) (string, error) {
    resp := &Qot_GetHoldingChangeList.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

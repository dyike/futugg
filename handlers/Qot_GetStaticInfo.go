package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_Common"
    "futugg/pb/Qot_GetStaticInfo"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3202), "Qot_GetStaticInfo")
    var err error
    err = futugg.On("send.Qot_GetStaticInfo", QotGetStaticInfoSend)
    err = futugg.On("recv.Qot_GetStaticInfo", QotGetStaticInfoRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetStaticInfoSend(conn *futugg.FutuGG, marketNum int32, secType int32, stockCode string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3202))

    var securityList []*Qot_Common.Security
    security := transStockCode(stockCode)
    securityList = append(securityList, security)

    reqData := &Qot_GetStaticInfo.Request{
        C2S: &Qot_GetStaticInfo.C2S{
            SecurityList: securityList,
            Market: &marketNum,
            SecType: &secType,
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

func QotGetStaticInfoRecv(data []byte) (string, error) {
    resp := &Qot_GetStaticInfo.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

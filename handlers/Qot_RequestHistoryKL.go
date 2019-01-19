package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_RequestHistoryKL"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3103), "Qot_RequestHistoryKL")
    var err error
    err = futugg.On("send.Qot_RequestHistoryKL", QotRequestHistoryKLSend)
    err = futugg.On("recv.Qot_RequestHistoryKL", QotRequestHistoryKLRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotRequestHistoryKLSend(conn *futugg.FutuGG, rehab string, kl string, stockCode string, beginTime string, endTime string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3103))

    rehabType := transRehabType(rehab)
    klType := transKLType(kl)
    security := transStockCode(stockCode)

    reqData := &Qot_RequestHistoryKL.Request{
        C2S: &Qot_RequestHistoryKL.C2S{
            RehabType: &rehabType,
            KlType: &klType,
            Security: security,
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

func QotRequestHistoryKLRecv(data []byte) (string, error) {
    resp := &Qot_RequestHistoryKL.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

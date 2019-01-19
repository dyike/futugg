package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_GetFunds"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2101), "Trd_GetFunds")
    var err error
    err = futugg.On("send.Trd_GetFunds", TrdGetFundsSend)
    err = futugg.On("recv.Trd_GetFunds", TrdGetFundsRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func TrdGetFundsSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2101))

    header := setTrdHeader(trdEnv, accID, trdMarket)
    reqData := &Trd_GetFunds.Request{
        C2S: &Trd_GetFunds.C2S{
            Header: header,
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

func TrdGetFundsRecv(data []byte) (string, error) {
    resp := &Trd_GetFunds.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

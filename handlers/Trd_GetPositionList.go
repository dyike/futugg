package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_GetPositionList"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(2102), "Trd_GetPositionList")
    var err error
    err = futugg.On("send.Trd_GetPositionList", Trd_GetPositionListSend)
    err = futugg.On("recv.Trd_GetPositionList", Trd_GetPositionListRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func Trd_GetPositionListSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32, ) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2102))

    header := setTrdHeader(trdEnv, accID, trdMarket)
    reqData := &Trd_GetPositionList.Request{
        C2S: &Trd_GetPositionList.C2S{
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

func Trd_GetPositionListRecv(data []byte) error {
    resp := &Trd_GetPositionList.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    fmt.Println(resp)

    return nil
}

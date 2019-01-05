package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_GetMaxTrdQtys"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(2111), "Trd_GetMaxTrdQtys")
    var err error
    err = futugg.On("send.Trd_GetMaxTrdQtys", TrdGetMaxTrdQtysSend)
    err = futugg.On("recv.Trd_GetMaxTrdQtys", TrdGetMaxTrdQtysRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func TrdGetMaxTrdQtysSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32, orderType int32, code string, price float64) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2111))

    header := setTrdHeader(trdEnv, accID, trdMarket)
    reqData := &Trd_GetMaxTrdQtys.Request{
        C2S: &Trd_GetMaxTrdQtys.C2S{
            Header: header,
            OrderType: &orderType,
            Code: &code,
            Price: &price,
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

func TrdGetMaxTrdQtysRecv(data []byte) error {
    resp := &Trd_GetMaxTrdQtys.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    fmt.Println(resp)

    return nil
}

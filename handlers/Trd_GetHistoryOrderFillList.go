package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_GetHistoryOrderFillList"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2222), "Trd_GetHistoryOrderFillList")
    var err error
    err = futugg.On("send.Trd_GetHistoryOrderFillList", TrdGetHistoryOrderFillListSend)
    err = futugg.On("recv.Trd_GetHistoryOrderFillList", TrdGetHistoryOrderFillListRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func TrdGetHistoryOrderFillListSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2222))

    header := setTrdHeader(trdEnv, accID, trdMarket)
    filterConditions := setTrdFilterConditions()
    reqData := &Trd_GetHistoryOrderFillList.Request{
        C2S: &Trd_GetHistoryOrderFillList.C2S{
            Header: header,
            FilterConditions: filterConditions,
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

func TrdGetHistoryOrderFillListRecv(data []byte) error {
    resp := &Trd_GetHistoryOrderFillList.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}

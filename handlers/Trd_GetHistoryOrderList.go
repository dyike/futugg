package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_GetHistoryOrderList"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2221), "Trd_GetHistoryOrderList")
    var err error
    err = futugg.On("send.Trd_GetHistoryOrderList", TrdGetHistoryOrderListSend)
    err = futugg.On("recv.Trd_GetHistoryOrderList", TrdGetHistoryOrderListRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func TrdGetHistoryOrderListSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2221))

    header := setTrdHeader(trdEnv, accID, trdMarket)
    filterConditions := setTrdFilterConditions()
    reqData := &Trd_GetHistoryOrderList.Request{
        C2S: &Trd_GetHistoryOrderList.C2S{
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

func TrdGetHistoryOrderListRecv(data []byte) (string, error) {
    resp := &Trd_GetHistoryOrderList.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_GetAccList"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2001), "Trd_GetAccList")
    var err error
    err = futugg.On("send.Trd_GetAccList", TrdGetAccListSend)
    err = futugg.On("recv.Trd_GetAccList", TrdGetAccListRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func TrdGetAccListSend(conn *futugg.FutuGG, userId uint64) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2001))

    reqData := &Trd_GetAccList.Request{
        C2S: &Trd_GetAccList.C2S{
            UserID: &userId,
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

func TrdGetAccListRecv(data []byte) error {
    resp := &Trd_GetAccList.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}

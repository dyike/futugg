package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Trd_SubAccPush"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2008), "Trd_SubAccPush")
    var err error
    err = futugg.On("send.Trd_SubAccPush", TrdSubAccPushSend)
    err = futugg.On("recv.Trd_SubAccPush", TrdSubAccPushRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func TrdSubAccPushSend(conn *futugg.FutuGG, accIDs string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2008))

    accIdList := transAccIDs(accIDs)
    reqData := &Trd_SubAccPush.Request{
        C2S: &Trd_SubAccPush.C2S{
            AccIDList: accIdList,
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

func TrdSubAccPushRecv(data []byte) (string, error) {
    resp := &Trd_SubAccPush.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

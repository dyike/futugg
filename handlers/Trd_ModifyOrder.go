package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Common"
    "futugg/pb/Trd_ModifyOrder"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2205), "Trd_ModifyOrder")
    var err error
    err = futugg.On("send.Trd_ModifyOrder", TrdModifyOrderSend)
    err = futugg.On("recv.Trd_ModifyOrder", TrdModifyOrderRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func TrdModifyOrderSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32, orderId uint64, modifyOrderOp int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2205))

    packetId := &Common.PacketID{
        ConnID: new(uint64),
        SerialNo: new(uint32),
    }
    *packetId.ConnID = conn.ConnID
    *packetId.SerialNo = pack.SeqNo

    header := setTrdHeader(trdEnv, accID, trdMarket)
    reqData := &Trd_ModifyOrder.Request{
        C2S: &Trd_ModifyOrder.C2S{
            PacketID: packetId,
            Header: header,
            OrderID: &orderId,
            ModifyOrderOp: &modifyOrderOp,
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

func TrdModifyOrderRecv(data []byte) error {
    resp := &Trd_ModifyOrder.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}

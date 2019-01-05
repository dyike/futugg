package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Common"
    "futugg/pb/Trd_PlaceOrder"

    "github.com/golang/protobuf/proto"
)

func init() {
    futugg.SetHandlerId(uint32(2201), "Trd_PlaceOrder")
    var err error
    err = futugg.On("send.Trd_PlaceOrder", TrdPlaceOrderSend)
    err = futugg.On("recv.Trd_PlaceOrder", TrdPlaceOrderRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func TrdPlaceOrderSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32, trdSide int32, orderType int32, code string, qty float64) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2201))

    packetId := &Common.PacketID{
        ConnID: new(uint64),
        SerialNo: new(uint32),
    }
    *packetId.ConnID = conn.ConnID
    *packetId.SerialNo = pack.SeqNo

    header := setTrdHeader(trdEnv, accID, trdMarket)
    reqData := &Trd_PlaceOrder.Request{
        C2S: &Trd_PlaceOrder.C2S{
            PacketID: packetId,
            Header: header,
            TrdSide: &trdSide,
            OrderType: &orderType,
            Code: &code,
            Qty: &qty,
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

func TrdPlaceOrderRecv(data []byte) error {
    resp := &Trd_PlaceOrder.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    fmt.Println(resp)

    return nil
}

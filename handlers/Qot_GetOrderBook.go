package handlers

import (
    "fmt"
    "futugg"
    "futugg/pb/Qot_GetOrderBook"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3012), "Qot_GetOrderBook")
    var err error
    err = futugg.On("send.Qot_GetOrderBook", QotGetOrderBookSend)
    err = futugg.On("recv.Qot_GetOrderBook", QotGetOrderBookRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetOrderBookSend(conn *futugg.FutuGG, stockCode string, num int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3012))

    security := transStockCode(stockCode)

    reqData := &Qot_GetOrderBook.Request{
        C2S: &Qot_GetOrderBook.C2S{
            Security: security,
            Num: &num,
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

func QotGetOrderBookRecv(data []byte) (string, error) {
    resp := &Qot_GetOrderBook.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return "", fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    return result, err
}

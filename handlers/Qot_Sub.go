package handlers

import (
	"fmt"
	"futugg"
	"futugg/pb/Qot_Common"
	"futugg/pb/Qot_Sub"

	"github.com/golang/protobuf/proto"
)

func init() {
	futugg.SetHandlerId(uint32(3001), "Qot_Sub")
	fmt.Println("112222")
	var err error
	err = futugg.On("send.Qot_Sub", QotSubSend)
	err = futugg.On("recv.Qot_Sub", QotSubRecv)
	if err != nil {
		fmt.Println(err)
	}
}

func QotSubSend(conn *futugg.FutuGG, stockCode string, subType string, isSubOrUnSub bool, isRegOrUnRegPush bool, isFirstPush bool) error {
	pack := &futugg.FutuPack{}
	pack.SetProto(uint32(3001))

	var securityList []*Qot_Common.Security
	security := transStockCode(stockCode)
	securityList = append(securityList, security)

	var subTypeList []int32
	subTypeNum := transSubType(subType)

	var regPushRehabTypeList []int32
	regPushRehabType := int32(1)
	regPushRehabTypeList = append(regPushRehabTypeList, regPushRehabType)

	reqData := &Qot_Sub.Request{
		C2S: &Qot_Sub.C2S{
			SecurityList:         securityList,
			SubTypeList:          subTypeList,
			IsSubOrUnSub:         &isSubOrUnSub,
			IsRegOrUnRegPush:     &isRegOrUnRegPush,
			RegPushRehabTypeList: regPushRehabTypeList,
			IsFirstPush:          &isFirstPush,
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

func QotSubRecv(data []byte) error {
	resp := &Qot_Sub.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(resp)

	return nil
}

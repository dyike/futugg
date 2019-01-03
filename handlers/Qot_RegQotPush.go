package handlers

import (
	"fmt"
	"futugg"
	"futugg/pb/Qot_Common"
	"futugg/pb/Qot_RegQotPush"

	"github.com/golang/protobuf/proto"
)

func init() {
	futugg.SetHandlerId(uint32(3002), "Qot_RegQotPush")
	var err error
	err = futugg.On("send.Qot_RegQotPush", QotRegQotPushSend)
	err = futugg.On("recv.Qot_RegQotPush", QotRegQotPushRecv)
	if err != nil {
		fmt.Println(err)
	}
}

func QotRegQotPushSend(conn *futugg.FutuGG, stockCode string, subType string, isRegOrUnReg bool, isFirstPush bool) error {
	pack := &futugg.FutuPack{}
	pack.SetProto(uint32(3002))

	var securityList []*Qot_Common.Security
	security := transStockCode(stockCode)
	securityList = append(securityList, security)

	var subTypeList []int32
	subTypeNum := transSubType(subType)
	subTypeList = append(subTypeList, subTypeNum)

	var regPushRehabTypeList []int32
	regPushRehabType := int32(1)
	regPushRehabTypeList = append(regPushRehabTypeList, regPushRehabType)

	reqData := &Qot_RegQotPush.Request{
		C2S: &Qot_RegQotPush.C2S{
			SecurityList:  securityList,
			SubTypeList:   subTypeList,
			IsRegOrUnReg:  &isRegOrUnReg,
			RehabTypeList: regPushRehabTypeList,
			IsFirstPush:   &isFirstPush,
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

func QotRegQotPushRecv(data []byte) error {
	resp := &Qot_RegQotPush.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(resp)

	return nil
}

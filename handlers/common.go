package handlers

import (
	"futugg/pb/Qot_Common"
	"futugg/pb/Trd_Common"
	"strings"
	"strconv"
	"fmt"
)

func transStockCode(code string) *Qot_Common.Security {
	codes := strings.Split(code, ".")
	stock := &Qot_Common.Security{
		Market: new(int32),
		Code:   new(string),
	}

	switch codes[0] {
	case "HK":
		*stock.Market = int32(1)
	case "US":
		*stock.Market = int32(11)
	case "SH":
		*stock.Market = int32(21)
	case "SZ":
		*stock.Market = int32(22)
	}

	*stock.Code = codes[1]
	return stock
}

func transAccIDs(accIds string) []uint64 {
	accIdsArr := strings.Split(accIds, ",")
	var accIDList []uint64
	for _, id := range accIdsArr {
		if s, err := strconv.ParseUint(id, 10, 64); err == nil {
			accIDList = append(accIDList, s)
		}
	}
	return accIDList
}

func setTrdFilterConditions() *Trd_Common.TrdFilterConditions {
	trdFilterConditions := &Trd_Common.TrdFilterConditions{
	}
	return trdFilterConditions
}

func setTrdHeader(trdEnv int32, accID uint64, trdMarket int32) *Trd_Common.TrdHeader {
	trdHeader := &Trd_Common.TrdHeader{
		TrdEnv: new(int32),
		AccID:  new(uint64),
		TrdMarket: new(int32),
	}
	*trdHeader.TrdEnv = trdEnv
	*trdHeader.AccID = accID
	*trdHeader.TrdMarket = trdMarket
	return trdHeader
}

func transSubType(subType string) int32 {
	k := "SubType_" + strings.Title(subType)
	if v, ok := Qot_Common.SubType_value[k]; ok {
		return v
	}

	return int32(0)
}

func transRehabType(rehabType string) int32 {
	k := "RehabType_" + strings.Title(rehabType)
	fmt.Println(k)
	if v, ok := Qot_Common.RehabType_value[k]; ok {
		return int32(v)
	}
	return int32(0)
}


func transKLType(KLType string) int32 {
	k := "KLType_" + strings.Title(KLType)
	if v, ok := Qot_Common.KLType_value[k]; ok {
		return v
	}
	// default KLType_Day = 2
	return int32(2)
}

package handlers

import (
	"futugg/pb/Qot_Common"
	"strings"
	"strconv"
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

func transSubType(subType string) int32 {
	k := "SubType_" + strings.Title(subType)
	if v, ok := Qot_Common.SubType_value[k]; ok {
		return v
	}

	return int32(0)
}

package handlers

import (
	"futugg/pb/Qot_Common"
	"strings"
)

func transStockCode(code string) *Qot_Common.Security {
	codes := strings.Split(code, ".")
	stock := &Qot_Common.Security{
		Market: new(int32),
		Code:   new(string),
	}

	switch codes[0] {
	case "HK":
		*stock.Marked = int32(1)
	case "US":
		*stock.Marked = int32(11)
	case "SH":
		*stock.Marked = int32(21)
	case "SZ":
		*stock.Marked = int32(22)
	}

	*stock.Code = codes[1]
	return stock
}

func transSubType(subType string) int32 {
	k := "SubType_" + strings.Title(subType)
	if v, ok := Qot_Common.SubType_value[k]; ok {
		return v
	}

	return int32(0)
}

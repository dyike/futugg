## Qot_GetOrderBook

获取买卖盘

### 示例

```go
futugg.Cmd("send.Qot_Sub", cli, "HK.01810", "OrderBook", true, true, "None", false)
futugg.Cmd("send.Qot_GetOrderBook", cli, "HK.01810", int32(2))
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
stockCode | string | 是 | 股票名称，固定格式 | HK.01810
num | int32 | 是 |请求的摆盘个数(1~10) | 2


### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "security": {
      "market": 1,
      "code": "01810"
    },
    "orderBookAskList": [   // 卖盘
      {
        "price": 0.0002654,    //委托价格
        "volume": "7172876309540896807",   //委托数量
        "orederCount": 2    //委托订单个数
      },
      {
        "price": 0.000442,
        "volume": "7086976963620896791",
        "orederCount": 2
      }
    ],
    "orderBookBidList": [  // 买盘
      {
        "price": 10.34,
        "volume": "32400",
        "orederCount": 6
      },
      {
        "price": 10.32,
        "volume": "1465600",
        "orederCount": 83
      }
    ]
  }
}
```



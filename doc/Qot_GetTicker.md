## Qot_GetTicker

获取逐笔

### 示例

```go
futugg.Cmd("send.Qot_Sub", cli, "HK.01810", "Ticker", true, true, "None", false)
futugg.Cmd("send.Qot_GetTicker", cli, "HK.01810", int32(2))
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
stockCode | string | 是 | 股票名称，固定格式 | HK.01810
maxRetNum | int32 | 是 |最多返回的逐笔个数,最多返回1000个 | 2


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
    "tickerList": [
      {
        "time": "2019-01-09 15:59:59",
        "sequence": "6644403737936804386",  // 唯一标识
        "dir": 2,     // TickerDirection, 买卖方向
        "price": 10.44,   
        "volume": "1000",
        "turnover": 10440,  
        "recvTime": 0,   //收到推送数据的本地时间戳，用于定位延迟
        "type": 1,    //TickerType, 逐笔类型
        "typeSign": 32    //逐笔类型符号
      },
      {
        "time": "2019-01-09 16:09:13",
        "sequence": "6644406117348686371",
        "dir": 3,
        "price": 10.34,
        "volume": "14349800",
        "turnover": 148376932,
        "recvTime": 0,
        "type": 7,
        "typeSign": 85
      }
    ]
  }
}
```



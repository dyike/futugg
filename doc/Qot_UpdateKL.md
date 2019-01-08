## Qot_UpdateKL

推送K线

### 示例

```go
futugg.Cmd("send.Qot_Sub", cli, "HK.01810", "KL_Month", true, true, "None", false)
futugg.Cmd("recv.Qot_UpdateKL", cli, "HK.01810")
```

### Request参数

不需要参数



### Response内容

```json 
{
  "retType": 0,
  "s2c": {
    "rehabType": 0,
    "klType": 4,
    "security": {
      "market": 1,
      "code": "01810"
    },
    "klList": [
      {
        "time": "2019-01-01 00:00:00",
        "isBlank": false,
        "highPrice": 12.98,
        "openPrice": 12.98,
        "lowPrice": 11.06,
        "closePrice": 11.1,
        "lastClosePrice": 0,
        "volume": "216853313",
        "turnover": 2548245258.7,
        "turnoverRate": 0.01382,
        "pe": 0,
        "changeRate": -14.483821263482286
      }
    ]
  }
}
```



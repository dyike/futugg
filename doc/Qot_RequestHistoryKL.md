## Qot_RequestHistoryKL

在线获取单只股票一段历史K线

### 示例

```go
futugg.Cmd("send.Qot_RequestHistoryKL", cli, "None", "1Min", "HK.01810", "2019-01-09 16:00:00", "2019-01-09 16:05:00")
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
rehab | string | 是 | 复权类型，None:不复权;Forward:前复权;Backward:后复权 | None
kl | string | 是 | K线类型, 1Min, Day... 具体参照下面的枚举类型 | 1Min
stockCode | string | 是 | 股票名称，固定格式 | HK.01810
beginTime | string | 是 | 开始时间 | 2019-01-09 16:00:00
endTime | string  | 是 | 结束时间 | 2019-01-09 16:05:00


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
    "klList": [
      {
        "time": "2019-01-09 16:00:00",
        "isBlank": false,
        "highPrice": 10.46,
        "openPrice": 10.42,
        "lowPrice": 10.34,
        "closePrice": 10.34,
        "lastClosePrice": 10.44,
        "volume": "15508600",
        "turnover": 160474764,
        "turnoverRate": 0,
        "pe": 0,
        "changeRate": -0.9578544061302648
      }
    ]
  }
}
```



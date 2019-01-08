## Qot_UpdateBasicQot

推送股票基本报价

### 示例

```go
futugg.Cmd("recv.Qot_UpdateBasicQot", cli)
```

### Request参数

不需要参数



### Response内容

```json 
{
  "retType": 0,
  "s2c": {
    "basicQotList": [
      {
        "security": {
          "market": 1,
          "code": "01810"
        },
        "isSuspended": false,
        "listTime": "2018-07-09",
        "priceSpread": 0.02,
        "updateTime": "2019-01-08 16:08:34",
        "highPrice": 12.16,
        "openPrice": 12.16,
        "lowPrice": 11.06,
        "curPrice": 11.1,
        "lastClosePrice": 12,
        "volume": "91106400",
        "turnover": 1034135901,
        "turnoverRate": 0.581,
        "amplitude": 9.167,
        "darkStatus": 0
      }
    ]
  }
}
```



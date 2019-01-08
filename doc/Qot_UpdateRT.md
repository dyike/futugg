## Qot_UpdateRT

推送分时

### 示例

```go
futugg.Cmd("send.Qot_Sub", cli, "HK.01810", "RT", true, true, "None", false)
futugg.Cmd("recv.Qot_UpdateRT", cli, "HK.01810")
```

### Request参数

不需要参数



### Response内容

```json 
{
  "retType": 0,
  "s2c": {
    "security": {
      "market": 1,
      "code": "01810"
    },
    "rtList": [
      {
        "time": "2019-01-08 16:00:00",
        "minute": 960,
        "isBlank": false,
        "price": 11.1,
        "lastClosePrice": 12,
        "avgPrice": 11.350255464035689,
        "volume": "2934600",
        "turnover": 32579188
      }
    ]
  }
}
```





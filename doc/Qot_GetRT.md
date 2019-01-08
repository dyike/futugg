## Qot_GetRT

获取分时

### 示例

```go
futugg.Cmd("send.Qot_Sub", cli, "HK.01810", "RT", true, true, "None", false)
futugg.Cmd("send.Qot_GetRT", cli, "HK.01810")
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
stockCode | string | 是 | 股票名称，固定格式 | HK.01810


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
    "rtList": [
      {
        "time": "2019-01-08 09:30:00",  //时间字符串
        "minute": 570,      //距离0点过了多少分钟 
        "isBlank": false,   //是否是空内容的点,若为ture则只有时间信息
        "price": 12.16,      //当前价
        "lastClosePrice": 12,    //昨收价
        "avgPrice": 12.16,    //均价
        "volume": "552800",  //成交量 
        "turnover": 6722048    //成交额
      },
      ...
      {
        "time": "2019-01-08 09:31:00",
        "minute": 571,
        "isBlank": false,
        "price": 11.82,
        "lastClosePrice": 12,
        "avgPrice": 11.999010343206393,
        "volume": "1148800",
        "turnover": 13695468
      }
    ]
  }
}
```



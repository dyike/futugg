## Qot_GetBasicQot

获取股票基本行情

### 示例

```go
futugg.Cmd("send.Qot_GetBasicQot", cli, "US.BILI")
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
stockCode | string | 是 | 股票名称，固定风格 [港股/美股].[股票编号] | 例如 HK.00700、US.BILI


### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "basicQotList": [
      {
        "security": {
          "market": 1,
          "code": "00700"
        },
        "isSuspended": false,
        "listTime": "2004-06-16",
        "priceSpread": 0.2,
        "updateTime": "2019-01-07 16:08:08",
        "highPrice": 318.8,
        "openPrice": 318,
        "lowPrice": 313.6,
        "curPrice": 317.6,
        "lastClosePrice": 310.6,
        "volume": "24890962",
        "turnover": 7881080851,
        "turnoverRate": 0.261,
        "amplitude": 1.674,
        "darkStatus": 0
      }
    ]
  }
}
```



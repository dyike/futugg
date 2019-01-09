## Qot_GetTradeDate

获取市场交易日

### 示例

```go
futugg.Cmd("send.Qot_GetTradeDate", cli, int32(1), "2019-01-08 16:00:00", "2019-01-09 16:00:00")
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
marketNum | int | 是 | 市场编号，如下QotMark枚举值 | 1
beginTime | string | 是 | 开始时间 | 2019-01-09 16:00:00
endTime | string  | 是 | 结束时间 | 2019-01-09 16:05:00

```
enum QotMarket
{
    QotMarket_Unknown = 0; //未知市场
    QotMarket_HK_Security = 1; //港股
    QotMarket_HK_Future = 2; //港期货(目前是恒指的当月、下月期货行情)
    QotMarket_US_Security = 11; //美股
    QotMarket_CNSH_Security = 21; //沪股
    QotMarket_CNSZ_Security = 22; //深股
}

```


### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "tradeDateList": [
      {
        "time": "2019-01-08"
      },
      {
        "time": "2019-01-09"
      }
    ]
  }
}
```



## Qot_GetSubInfo

获取订阅信息

### 示例

```go
futugg.Cmd("send.Qot_Sub", cli, "HK.01810", "KL_1Min", true, true, "None", false)
futugg.Cmd("send.Qot_GetKL", cli, "None", "1Min", "HK.01810", 1)
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
rehab | string | 是 | 复权类型，None:不复权;Forward:前复权;Backward:后复权 | None
kl | string | 是 | K线类型, 1Min, Day... 具体参照下面的枚举类型 | 1Min
stockCode | string | 是 | 股票名称，固定格式 | HK.01810
reqNum | int32 | 是 | 请求K线根数 | 1

```
//枚举值兼容旧协议定义
//新类型季K,年K,3分K暂时没有支持历史K线
enum KLType
{
        KLType_1Min = 1; //1分K
        KLType_Day = 2; //日K
        KLType_Week = 3; //周K
        KLType_Month = 4; //月K
        KLType_Year = 5; //年K
        KLType_5Min = 6; //5分K
        KLType_15Min = 7; //15分K
        KLType_30Min = 8; //30分K
        KLType_60Min = 9; //60分K
        KLType_3Min = 10; //3分K
        KLType_Quarter = 11; //季K
}

```


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
        "time": "2019-01-08 16:00:00",     //时间戳字符串
        "isBlank": false,                 //是否是空内容的点,若为ture则只有时间信息
        "highPrice": 11.14,               //最高价
        "openPrice": 11.12,               //开盘价
        "lowPrice": 11.1,                 //最低价
        "closePrice": 11.1,               //收盘价
        "lastClosePrice": 11.12,          //昨收价
        "volume": "2934600",              //成交量
        "turnover": 32579188,             //成交额
        "turnoverRate": 0,                //换手率
        "pe": 0,                          //市盈率
        "changeRate": -0.17985611510790983 //涨跌幅
      }
    ]
  }
}
```



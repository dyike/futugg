## Qot_GetStaticInfo

获取股票静态信息

### 示例

```go
futugg.Cmd("send.Qot_GetStaticInfo", cli, int32(1), int32(3), "HK.01810")
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
marketNum | int | 是 | 市场编号，如下QotMark枚举值 | 1
secType | int | 是 | 证券类型，如下SecurityType枚举值 | 3
stockCode | string | 是 | 股票标号 | "HK.01810"

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

enum SecurityType
{
        SecurityType_Unknown = 0; //未知
        SecurityType_Bond = 1; //债券
        SecurityType_Bwrt = 2; //一揽子权证
        SecurityType_Eqty = 3; //正股
        SecurityType_Trust = 4; //信托,基金
        SecurityType_Warrant = 5; //涡轮
        SecurityType_Index = 6; //指数
        SecurityType_Plate = 7; //板块
        SecurityType_Drvt = 8; //期权
        SecurityType_PlateSet = 9; //板块集
}

```


### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "staticInfoList": [
      {
        "basic": {
          "security": {
            "market": 1,
            "code": "01810"
          },
          "id": "76033806042898",  //股票ID
          "lotSize": 200,   //每手数量
          "secType": 3,     //Qot_Common.SecurityType,股票类型
          "name": "小米集团-W",
          "listTime": "2018-07-09",
          "delisting": false
        }
      }
    ]
  }
}
```



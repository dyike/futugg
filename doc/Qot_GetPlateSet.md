## Qot_GetPlateSet

获取板块集合下的板块

### 示例

```go
futugg.Cmd("send.Qot_GetPlateSet", cli, int32(1), int32(0))
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
marketNum | int | 是 | 市场编号，如下QotMark枚举值 | 1
plateSetType | int | 是 | 板块集合类型，如下PlateSetType枚举值 | 0

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

enum PlateSetType
{
        PlateSetType_All = 0; //所有板块
        PlateSetType_Industry = 1; //行业板块
        PlateSetType_Region = 2; //地域板块,港美股市场的地域分类数据暂为空
        PlateSetType_Concept = 3; //概念板块
        PlateSetType_Other = 4; //其他板块, 仅用于3207（获取股票所属板块）协议返回,不可作为其他协议的请求参数
}

```


### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "plateInfoList":[
      {
        "plate": {
          "market": 1,
          "code": "BK1175"
        },
        "name": "一带一路"
      },
      {
        "plate": {
          "market": 1,
          "code": "BK1922"
        },
        "name": "港股通(深)"
      },
      {
        "plate": {
          "market": 1,
          "code": ".FTIHK"
        },
        "name": "富途港股新经济指数"
      }
    ]
  }
}
```



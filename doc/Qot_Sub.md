## Qot_Sub

订阅或者反订阅

### 示例

```go
futugg.Cmd("send.Qot_Sub", cli, "US.BILI", "Basic", true, false, true)
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
stockCode | string | 是 | 股票名称，固定风格 [港股/美股].[股票编号] | 例如 HK.00700、US.BILI
subType | string | 是 | 订阅数据类型 | None/Basic，只需要SubType_前缀后面的内容
isSubOrUnSub | bool | 是 | ture表示订阅,false表示反订阅 | true
isRegOrUnRegPush | bool | 否 | 注册或反注册该连接上面行情的推送 | 默认false，不做注册反注册操作
isFirstPush | bool | 否 | 注册后如果本地已有数据是否首推一次已存在数据 | 默认 true 


关于`subType`字段的选择：

```
enum SubType
{
        SubType_None = 0;
        SubType_Basic = 1; //基础报价
        SubType_OrderBook = 2; //摆盘
        SubType_Ticker = 4; //逐笔
        SubType_RT = 5; //分时
        SubType_KL_Day = 6; //日K
        SubType_KL_5Min = 7; //5分K
        SubType_KL_15Min = 8; //15分K
        SubType_KL_30Min = 9; //30分K
        SubType_KL_60Min = 10; //60分K
        SubType_KL_1Min = 11; //1分K
        SubType_KL_Week = 12; //周K
        SubType_KL_Month = 13; //月K
        SubType_Broker = 14; //经纪队列
        SubType_KL_Qurater = 15; //季K
        SubType_KL_Year = 16; //年K
        SubType_KL_3Min = 17; //3分K
}
```


### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0
}
```



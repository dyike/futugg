## GetGlobalState

获取全局状态

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
userID | uint64 | 是 | 用户ID，必须跟FutuOpenD登陆的牛牛用户ID一致，否则会返回失败 | 123456

在futugg中，发送`GetGlobalState`请求，不需要额外传入参数，已经根据InitConnect实现中取得userID，传入其中。

### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "marketHK": 6,    
    "marketUS": 11,
    "marketSH": 6,
    "marketSZ": 6,
    "marketHKFuture": 14,
    "qotLogined": true,
    "trdLogined": true,
    "serverVer": 108,
    "serverBuildNo": 266,
    "time": "1546790498",
    "localTime": 1546790498.521573
  }
}
```



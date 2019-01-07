## Qot_GetSubInfo

获取订阅信息

### 示例

```go
futugg.Cmd("send.Qot_GetSubInfo", cli, true)
```

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
isReqAllConn | bool | 是 | 所有连接的订阅状态，所有连接的订阅状态 | true


### Response内容

```json 
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "connSubInfoList": [
      {
        "subInfoList": [    // 该连接订阅信息 
          {
            "subType": 1,
            "securityList": [
              {
                "market": 11,
                "code": "BILI"
              }
            ]
          }
        ],
        "usedQuota": 1,    //该连接已经使用的订阅额度
        "isOwnConnData": false   // 用于区分是否是自己连接的数据
      }
    ],
    "totalUsedQuota": 1,   // FutuOpenD已使用的订阅额度
    "remainQuota": 299     // FutuOpenD剩余订阅额度
  }
}
```



## InitConnect

初始化连接，请求中所有的协议必须等InitConnect协议先完成，在futugg.go文件中可以看到initConnect()的实现。

### Request参数

参数名  | 类型  | 是否必须 | 说明      | 示例
------- | ---- | -------- | -------  | ---------
clientVer | int32 | 是 | 客户端版本号，clientVer = "."以前的数 * 100 + "."以后的。 | 1.1版本的clientVer为1 * 100 + 1 = 101，2.21版本为2 * 100 + 21 = 221
clientID | string | 是 |客户端唯一标识，无生具体生成规则，客户端自己保证唯一性即可 | futugg_1234
recvNotify | bool | 否 | 此连接是否接收市场状态、交易需要重新解锁等等事件通知,true代表接收，反之不接收 | true

### Response内容

```json
{
  "retType": 0,
  "retMsg": "",
  "errCode": 0,
  "s2c": {
    "serverVer": 108,
    "loginUserID": "your userID",
    "connID": "连接ID，唯一标识",
    "connAESKey": "AES加密key，规定16字节长字符串",
    "keepAliveInterval": 10
  }
}

```

如果配置了加密，则需要使用`connAESKey`作为协议加密key。

`keepAliveInterval`为建议client发起的心跳KeepAlive间隔。SDK中已经实现默认为3s发一次KeepAlive的请求。

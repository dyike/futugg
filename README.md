## FutuGG

FutuGG是基于FutuOpenD网关实现的Golang的SDK包，项目中使用了原生了proto文件，注意不要更改生成的go文件，修改了import路径。

### 特别声明

因为futugg是跟金钱打交道的，可能你会用来做交易。如果在使用过程中触发了某个bug，导致金钱损失，一切后果请自行承担！！！

### 环境依赖

* 安装Google的protobuf，protoc-gen-go工具
* Golang1.9及最新版本

### 使用方式

```go

package main

import (
    "fmt"
    "futugg"
    _ "futugg/handlers"
)

func main() {
    block := make(chan bool)
    cli := futugg.New("0.0.0.0", "11250", "")

    futugg.Cmd("send.GetGlobalState", cli)
    cli.KeepAlive()

    // recv
    go func() {
        fmt.Println("start recv data")
        cli.Recv()
    }()

    futugg.Cmd("send.Qot_GetSubInfo", cli, true)
    // 3001 Qot_Sub
    futugg.Cmd("send.Qot_Sub", cli, "SZ.300104", "Basic", true, false, "None", false)

    <-block
}

```


使用前，需要启动富途网关OpenD，具体操作详见[futu-api-doc介绍](https://futunnopen.github.io/futu-api-doc/intro/FutuOpenDGuide.html)

支持同步和异步两种模式


### 广告

如有需要，可以填写推荐人牛牛号：7189075，谢谢哟！

### 文档简介

ProtoId | 名称 | 说明
------- | ---- | ----------
1001    | [InitConnect](https://github.com/dyike/futugg/blob/master/doc/InitConnect.md)   |       初始化连接
1002    | [GetGlobalState](https://github.com/dyike/futugg/blob/master/doc/GetGlobalState.md)    |    获取全局状态
1003    | Notify    |    系统通知推送
1004    | KeepAlive    | 保活心跳
2001    | Trd_GetAccList    |    获取业务账户列表
2005    | Trd_UnlockTrade    |   解锁或锁定交易
2008    | Trd_SubAccPush    |    订阅业务账户的交易推送数据
2101    | Trd_GetFunds    |  获取账户资金
2102    | Trd_GetPositionList    |   获取账户持仓
2111    | Trd_GetMaxTrdQtys    | 获取最大交易数量
2201    | Trd_GetOrderList    |  获取订单列表
2202    | Trd_PlaceOrder    |    下单
2205    | Trd_ModifyOrder    |   修改订单
2208    | Trd_UpdateOrder    |   推送订单状态变动通知
2211    | Trd_GetOrderFillList    |  获取成交列表
2218    | Trd_UpdateOrderFill    |   推送成交通知
2221    | Trd_GetHistoryOrderList    |   获取历史订单列表
2222    | Trd_GetHistoryOrderFillList    |   获取历史成交列表
3001    | [Qot_Sub](https://github.com/dyike/futugg/blob/master/doc/Qot_Sub.md)      |   订阅或者反订阅
3002    | [Qot_RegQotPush](https://github.com/dyike/futugg/blob/master/doc/Qot_RegQotPush.md)    |    注册推送
3003    | [Qot_GetSubInfo](https://github.com/dyike/futugg/blob/master/doc/Qot_GetSubInfo.md)     |    获取订阅信息
3004    | [Qot_GetBasicQot](https://github.com/dyike/futugg/blob/master/doc/Qot_GetBasicQot.md)    |   获取股票基本行情
3005    | [Qot_UpdateBasicQot](https://github.com/dyike/futugg/blob/master/doc/Qot_UpdateBasicQot.md)      |    推送股票基本行情
3006    | [Qot_GetKL](https://github.com/dyike/futugg/blob/master/doc/Qot_GetKL.md)      | 获取K线
3007    | [Qot_UpdateKL](https://github.com/dyike/futugg/blob/master/doc/Qot_UpdateKL.md)     |  推送K线
3008    | [Qot_GetRT](https://github.com/dyike/futugg/blob/master/doc/Qot_GetRT.md)    | 获取分时
3009    | [Qot_UpdateRT](https://github.com/dyike/futugg/blob/master/doc/Qot_UpdateRT.md)    |  推送分时
3010    | [Qot_GetTicker](https://github.com/dyike/futugg/blob/master/doc/Qot_GetTicker.md)    | 获取逐笔
3011    | Qot_UpdateTicker    |  推送逐笔
3012    | [Qot_GetOrderBook](https://github.com/dyike/futugg/blob/master/doc/Qot_GetOrderBook.md)    |  获取买卖盘
3013    | Qot_UpdateOrderBook    |   推送买卖盘
3014    | [Qot_GetBroker](https://github.com/dyike/futugg/blob/master/doc/Qot_GetBroker.md)     | 获取经纪队列
3015    | Qot_UpdateBroker    |  推送经纪队列
~~3100~~| ~~Qot_GetHistoryKL~~|  ~~从本地下载历史数据获取单只股票一段历史K线~~
~~3101~~| ~~Qot_GetHistoryKLPoints~~|    ~~从本地下载历史数据获取多只股票多点历史K线~~
~~3102~~| ~~Qot_GetRehab~~    |  ~~从本地下载历史数据获取复权信息~~
3103    | Qot_RequestHistoryKL    |  在线获取单只股票一段历史K线
3200    | Qot_GetTradeDate    |  获取市场交易日
3202    | Qot_GetStaticInfo    | 获取股票静态信息
3203    | Qot_GetSecuritySnapshot    |   获取股票快照
3204    | Qot_GetPlateSet    |   获取板块集合下的板块
3205    | Qot_GetPlateSecurity    |  获取板块下的股票
3206    | Qot_GetReference    |  获取正股相关股票
3207    | Qot_GetOwnerPlate    | 获取股票所属板块
3208    | Qot_GetHoldingChangeList    |  获取持股变化列表














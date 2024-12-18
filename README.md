# Cusdt ( Easy Payment Usdt)

<p align="center">
<img src="./static/img/tether.svg" width="15%" alt="tether">
</p>
<p align="center">
<a href="https://www.gnu.org/licenses/gpl-3.0.html"><img src="https://img.shields.io/badge/license-GPLV3-blue" alt="license GPLV3"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/Golang-1.22-red" alt="Go version 1.21"></a>
<a href="https://github.com/gin-gonic/gin"><img src="https://img.shields.io/badge/Gin-v1.9-blue" alt="Gin Web Framework v1.9"></a>
<a href="https://github.com/go-telegram-bot-api/telegram-bot-api"><img src="https://img.shields.io/badge/Telegram Bot-v5-lightgrey" alt="Golang Telegram Bot Api-v5"></a>
<a href="https://github.com/v03413/bepusdt"><img src="https://img.shields.io/badge/Release-v1.9.21-green" alt="Release v1.9.21"></a>
</p>


## 🪧 介绍

基本就是对`Epusdt`重新造了一次轮子，移除一些非必要依赖(`Redis MySQL`)，同时加入一些新特性，让个人`USDT.TRC20`
收款更好用、部署更便捷！

## 🎉 新特性

- ✅ 具备`Epusdt`的所有特性，插件兼容无缝替换
- ✅ USDT汇率实时同步交易所，且支持在以此基础上波动
- ✅ 不依赖`MySQL Redis`环境，支持`Docker`部署
- ✅ 支持非订单交易监控通知，钱包余额变动及时通知
- ✅ 机器人支持查询当前实时汇率，计算实际浮动汇率
- ✅ 机器人支持任意地址查询 USDT、TRX余额等信息
- ✅ 订单收款成功和余额变动通知 支持指定群组推送
- ✅ 支持可自定义USDT支付数额的精度，递增颗粒度
- ✅ 底层直接采用区块扫描，不依赖三方API，秒级响应
- ✅ <u>支持【TRX】收款</u>， TRX余额变动监控通知
- ✅ 支持钱包地址 <u>能量代理</u>和<u>能量回收</u> 监控通知

### 参数列表

| 参数名称                 | 默认值                  | 用法说明                                                                                                                                          |
|----------------------|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| EXPIRE_TIME          | `600`                | 订单有效期，单位秒                                                                                                                                     |
| USDT_RATE            | `6.4`                | USDT汇率，默认留空则获取Okx交易所的汇率(每分钟同步一次)，支持多种写法，如：`7.4` 表示固定7.4、`～1.02`表示最新汇率上浮2%、`～0.97`表示最新汇率下浮3%、`+0.3`表示最新加0.3、`-0.2`表示最新减0.2，以此类推；如参数错误则使用固定值6.4 |
| USDT_ATOM            | `0.01`               | `0.01`表示支付数额保留两位小数，相同金额时递增颗粒度为`0.01`，依次类推，如无特殊需求不建议修改。                                                                                        |
| TRX_RATE             | `0.95`               | **同 USDT_RATE**                                                                                                                               |
| TRX_ATOM             | `0.01`               | **同 USDT_ATOM**                                                                                                                               |
| AUTH_TOKEN           | `123234`             | 认证Token，对接会用到这个参数                                                                                                                             |
| LISTEN               | `:8080`              | 服务器HTTP监听地址                                                                                                                                   |
| TRADE_IS_CONFIRMED   | `0`                  | 是否需要网络确认，禁用可以提高回调速度，启用则可以防止交易失败                                                                                                               |
| APP_URI              | 空                    | 应用访问地址，留空则系统自动获取，前端收银台会用到，建议设置，例如：https://token-pay.example.com                                                                               |
| WALLET_ADDRESS       | 空                    | 启动时需要添加的钱包地址，多个请用半角符逗号`,`分开；当然，同样也支持通过机器人添加。                                                                                                  |
| TG_BOT_TOKEN         | 空                    | Telegram Bot Token，**必须设置**，否则无法使用                                                                                                            |
| TG_BOT_ADMIN_ID      | 空                    | Telegram Bot 管理员ID，**必须设置**，否则无法使用                                                                                                            |
| TG_BOT_GROUP_ID      | 空                    | Telegram 群组ID，设置之后机器人会将交易消息会推送到此群                                                                                                             |
| PAYMENT_AMOUNT_RANGE | `0.01,99999`         | 支付监控的允许数额范围(闭区间)，设置合理数值可避免一些诱导式诈骗交易提醒                                                                                                         |
| TRON_GRPC_NODE       | `18.141.79.38:50051` | Tron区块网络的GRPC节点，可选列表：https://developers.tron.network/docs/networks#public-node                                                                |
| STATIC_PATH          | 空                    | 静态资源路径，例如：`/root/bepusdt/static`；通过此参数可自定义模板，参考`static`目录，如非必要不建议修改。                                                                          |

**Ps：所以综上所述，必须设置的参数有`TG_BOT_TOKEN TG_BOT_ADMIN_ID`，否则无法使用！**

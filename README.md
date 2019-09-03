# telegram 关键词自动回复机器人 开源版本
这个是关键词回复机器人的开源版本代码，[Release](https://github.com/zu1k/tg-keyword-reply-bot/releases) 中发布的为后续闭源版本，功能更多。

## 开源版本介绍
### 基本命令介绍
- 添加关键词回复规则 `/add 关键词===回复内容` 或者 `/add 关键词1||关键词2===回复内容` 
- 关键词可以使用正则表达式,例如`/add re:p([a-z]+)ch===测试正则`,就会匹配规则`p([a-z]+)ch`  
- 删除关键词规则 `/del 关键词` 暂不支持一次性删除多个关键词
- 自动删除含有关键词的文字消息, 只需要将回复内容设置成 `delete`, 并给机器人添加删除消息权限
- 使用`/list`命令可以查看本群内所有自动回复规则
- 给机器人添加删除消息和踢人的管理权限,可以自动防清真(阿拉伯语)

### 回复特殊内容
- 回复内容支持文字\图片\GIF\视频,默认文字
- 如需图片,回复内容设置成`photo:https://t.me/c/1472018167/53095`,`https://t.me/c/1472018167/53095`是已经发送过的图片获取到的链接
- 同理,gif将`photo`替换成`gif`,视频替换成`video`,文件替换成`file`
- 注意: 这里的链接必须是公开群组的,否则无法发出来


## 闭源版本介绍
机器人闭源版本不断更新，至今已增加更多功能，详细功能见 **[使用说明](https://telegra.ph/%E8%BF%99%E4%B8%AA%E6%88%91%E7%9F%A5%E9%81%93%E6%9C%BA%E5%99%A8%E4%BA%BA%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E-07-07)**

### 使用闭源版本可执行文件搭建  
进入 **[Release 页面](https://github.com/zu1k/tg-keyword-reply-bot/releases)** 下载最新版本可执行文件到服务器，然后进行部署       
- 系统推荐使用： Ubuntu 18.04, glibc版本2.27
- 初次使用 `./tg-keyword-reply tg-bot-token` , 会将token存到数据库中
- 后面使用 `./tg-keyword-reply` , 无需输入token

#### 服务器选择
因为telegram bot api服务器在荷兰，所以搭建机器人推荐荷兰服务器，比较好的推荐vultr      
你可以选择使用我的aff链接注册vultr: https://www.vultr.com/?ref=7791688-4F      
因为机器人使用golang编写，性能较高，无需购买高配服务器，选择vultr最低配1核1G足以

#### 机器人命令列表
是推荐给用户显示的命令，私聊机器人爹设置

```
help - 查看帮助
add - 添加规则
del - 删除规则
list - 列出规则
admin - 呼叫管理员
banme - 禁言小游戏
getid - 查看用户的信息 可回复查看别人
autoreply - 开关自动回复功能
autodelete - 开关自动删除消息功能
replyorder - 开关回复ban/kick命令功能
banmegame - 开关禁言小游戏功能
playorderban - 开关玩命令惩罚功能
banqingzhen - 开关防清真功能
calladmin - 开关呼叫管理员功能
welcome - 开关加群欢迎功能
goodbye - 开关离群送别功能
deletejoinmessage - 开关删除加群消息功能
servicelist - 查看机器人功能列表
```
#### 注意事项
因为业务需要，需要私聊机器人爹关闭 Privacy Mode

## 问题反馈
欢迎大家反馈问题，在阅读了 **[使用说明](https://telegra.ph/%E8%BF%99%E4%B8%AA%E6%88%91%E7%9F%A5%E9%81%93%E6%9C%BA%E5%99%A8%E4%BA%BA%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E-07-07)** 之后仍旧无法解决问题，可以通过下满两种方式反馈
- 进入 [反馈群](https://t.me/keywordreplybotcallback) 反馈问题 （推荐，实时性强）
- 在github本库Issue反馈 （不太建议）


## 目前在线机器人
如果有同学自己搭建的，也可以私聊我添加进列表
- [这个我知道](https://t.me/keyword_reply_bot)  我搭建
- [关键词机器人](https://t.me/keywordreplybot)  我搭建
- [飞行中国](https://t.me/WeedChina_bot)  [农夫 420](https://t.me/nongfu420)搭建


## 赞助
指 [这个我知道](https://t.me/keyword_reply_bot) 的服务器费用赞助，目前机器运行在一台1核2G的服务器上，月费用10刀，希望有 **钞能力** 的同学赞助一下服务器费用。

### 赞助方式
- tg私聊我： [Cop](https://t.me/veezer) 进行赞助（无手续费)
- 进入发卡平台购买虚拟卡片进行赞助： [发卡平台](https://www.kuaifaka.com/purchasing?link=peekfun) (平台收取一定手续费)

### 目前赞助名单
- [聪聪](https://t.me/congcong) 420元
- [小明HR](https://t.me/xuezha) 36元
- [阿雅](https://t.me/alin0524) 50元 


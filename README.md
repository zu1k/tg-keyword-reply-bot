# Telegram 关键词自动回复机器人
<p align="center">
    <a href="https://goreportcard.com/report/github.com/zu1k/tg-keyword-reply-bot">
        <img src="https://goreportcard.com/badge/github.com/zu1k/tg-keyword-reply-bot">
    </a>
    <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Fzu1k%2Ftg-keyword-reply-bot?ref=badge_shield">
        <img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Fzu1k%2Ftg-keyword-reply-bot.svg?type=shield" alt="FOSSA Status">
    </a>
</p>

关键词机器人分开源和闭源两个版本，开源版本包含最基础的关键词回复功能，闭源版本增加了更多实用功能。                 
本项目为关键词自动回复机器人的开源代码，在 [Release](https://github.com/zu1k/tg-keyword-reply-bot/releases) 中发布的为闭源版本可执行文件。

## 开源版本
### 基本命令
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


## 闭源版本
闭源版本增加更多实用功能，可执行文件见 [Release](https://github.com/zu1k/tg-keyword-reply-bot/releases) 页面。

### 机器人命令列表
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

## 使用说明
使用说明见 [博客](https://blog.lgf.im/2019/11/telegram-keyword-reply-bot.html)

## 在线机器人
- [这个我知道](https://t.me/keyword_reply_bot)  拒绝博彩、狗推、洗钱等群组                 
如需付费搭建独享无限制机器人，麻烦邮箱联系我


## 赞助
赞助并非赞赏，所有费用皆用于服务器支付。           
进入发卡平台购买虚拟卡片进行赞助： [发卡平台](https://www.kuaifaka.com/purchasing?link=peekfun)      

### 赞助名单（感谢）
- [聪聪](https://t.me/congcong) 420元+860元+1664元
- [小明HR](https://t.me/xuezha) 36元
- [阿雅](https://t.me/alin0524) 50元 
- [冠希 科技传媒](https://t.me/a12399999) 39元
- [🆉🄴🄰🄻🅂🄾🄽](https://t.me/zealson) 50元
- [古博VPS](https://t.me/guboorg) 200元
- [LaN](https://t.me/BGdfd) 140元

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fzu1k%2Ftg-keyword-reply-bot.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fzu1k%2Ftg-keyword-reply-bot?ref=badge_large)

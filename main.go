package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
)

var bot *api.BotAPI
var gcron *cron.Cron

func main() {
	botToken := flag.String("t", "", "your bot Token")
	flag.IntVar(&superUserId, "s", 0, "super manager Id")
	flag.Parse()
	token := dbInit(*botToken)
	gcron = cron.New()
	gcron.Start()
	//开始工作
	start(token)
}

func start(botToken string) {
	bott, err := api.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	bot = bott
	bot.Debug = true
	log.Printf("Authorized on account: %s  ID: %d", bot.Self.UserName, bot.Self.ID)

	u := api.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic("Can't get Updates")
	}
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		go processUpdate(update)
	}
}

/**
 * 对于每一个update的单独处理
 */
func processUpdate(update api.Update) {
	var msg api.MessageConfig
	upmsg := update.Message
	gid := upmsg.Chat.ID
	uid := upmsg.From.ID
	//检查是不是新加的群或者新开的人
	in := checkInGroup(gid)
	if !in { //不在就需要加入, 内存中加一份, 数据库中添加一条空规则记录
		fmt.Println("新群組：", gid)
		groups = append(groups, gid)
		allkvs[gid] = make(Kvs)
		dbAddNewGroup(gid)
	}
	if upmsg.IsCommand() {
		msg = api.NewMessage(gid, "")
		_, _ = bot.DeleteMessage(api.NewDeleteMessage(gid, upmsg.MessageID))
		switch upmsg.Command() {
		case "start", "help":
			msg.Text = "本机器人能够自动回复特定关键词"
			sendMessage(msg)
		case "add":
			if checkAdmin(gid, *upmsg.From) {
				order := upmsg.CommandArguments()
				if order != "" {
					addRule(gid, order)
					msg.Text = "规则添加成功: " + order
				} else {
					msg.Text = addText
					msg.ParseMode = "Markdown"
					msg.DisableWebPagePreview = true
				}
				sendMessage(msg)
			}
		case "del":
			if checkAdmin(gid, *upmsg.From) {
				order := upmsg.CommandArguments()
				if order != "" {
					delRule(gid, order)
					msg.Text = "规则删除成功: " + order
				} else {
					msg.Text = delText
					msg.ParseMode = "Markdown"
				}
				sendMessage(msg)
			}
		case "list":
			if checkAdmin(gid, *upmsg.From) {
				rulelists := getRuleList(gid)
				msg.Text = "ID: " + strconv.FormatInt(gid, 10)
				msg.ParseMode = "Markdown"
				msg.DisableWebPagePreview = true
				sendMessage(msg)
				for _, rlist := range rulelists {
					msg.Text = rlist
					msg.ParseMode = "Markdown"
					msg.DisableWebPagePreview = true
					sendMessage(msg)
				}
			}
		case "admin":
			msg.Text = "[" + upmsg.From.String() + "](tg://user?id=" + strconv.Itoa(uid) + ") 请求管理员出来打屁股\r\n\r\n" + getAdmins(gid)
			msg.ParseMode = "Markdown"
			sendMessage(msg)
			banMember(gid, uid, 30)
		case "banme":
			botme, _ := bot.GetChatMember(api.ChatConfigWithUser{ChatID: gid, UserID: bot.Self.ID})
			if botme.CanRestrictMembers {
				rand.Seed(time.Now().UnixNano())
				sec := rand.Intn(540) + 60
				banMember(gid, uid, int64(sec))
				msg.Text = "恭喜[" + upmsg.From.String() + "](tg://user?id=" + strconv.Itoa(upmsg.From.ID) + ")获得" + strconv.Itoa(sec) + "秒的禁言礼包"
				msg.ParseMode = "Markdown"
			} else {
				msg.Text = "请给我禁言权限,否则无法进行游戏"
			}
			sendMessage(msg)
		case "me":
			myuser := upmsg.From
			msg.Text = "[" + upmsg.From.String() + "](tg://user?id=" + strconv.Itoa(upmsg.From.ID) + ") 的账号信息" +
				"\r\nID: " + strconv.Itoa(uid) +
				"\r\nUseName: [" + upmsg.From.String() + "](tg://user?id=" + strconv.Itoa(upmsg.From.ID) + ")" +
				"\r\nLastName: " + myuser.LastName +
				"\r\nFirstName: " + myuser.FirstName +
				"\r\nIsBot: " + strconv.FormatBool(myuser.IsBot)
			msg.ParseMode = "Markdown"
			sendMessage(msg)
		default:
		}
	} else {
		//回复类型的管理命令
		if upmsg.ReplyToMessage != nil {
			reply_to_memid := upmsg.ReplyToMessage.From.ID
			switch upmsg.Text {
			case "ban":
				if checkAdmin(gid, *upmsg.From) {
					banMember(gid, reply_to_memid, -1)
					mem, _ := bot.GetChatMember(api.ChatConfigWithUser{ChatID: gid, SuperGroupUsername: "", UserID: reply_to_memid})
					if !mem.CanSendMessages {
						msg = api.NewMessage(gid, "")
						msg.Text = "[" + upmsg.From.String() + "](tg://user?id=" + strconv.Itoa(upmsg.From.ID) + ") 禁言了 " +
							"[" + upmsg.ReplyToMessage.From.String() + "](tg://user?id=" + strconv.Itoa(reply_to_memid) + ") "
						msg.ParseMode = "Markdown"
						sendMessage(msg)
					}
				}
			case "unban":
				if checkAdmin(gid, *upmsg.From) {
					unbanMember(gid, reply_to_memid)
					//mem,_ := bot.GetChatMember(api.ChatConfigWithUser{gid, "", reply_to_memid})
					//
					msg = api.NewMessage(gid, "")
					msg.Text = "[" + upmsg.From.String() + "](tg://user?id=" + strconv.Itoa(upmsg.From.ID) + ") 解禁了 " +
						"[" + upmsg.ReplyToMessage.From.String() + "](tg://user?id=" + strconv.Itoa(reply_to_memid) + ") "
					msg.ParseMode = "Markdown"
					sendMessage(msg)
				}
			case "kick":
				if checkAdmin(gid, *upmsg.From) {
					kickMember(gid, reply_to_memid)
				}
			case "unkick":
				if checkAdmin(gid, *upmsg.From) {
					unkickMember(gid, reply_to_memid)
				}
			default:
			}
		}

		replyText := findKey(gid, upmsg.Text)
		if replyText == "delete" {
			_, _ = bot.DeleteMessage(api.NewDeleteMessage(gid, upmsg.MessageID))
		} else if strings.HasPrefix(replyText, "ban") {
			_, _ = bot.DeleteMessage(api.NewDeleteMessage(gid, upmsg.MessageID))
			banMember(gid, uid, -1)
		} else if strings.HasPrefix(replyText, "kick") {
			_, _ = bot.DeleteMessage(api.NewDeleteMessage(gid, upmsg.MessageID))
			kickMember(gid, uid)
		} else if strings.HasPrefix(replyText, "photo:") {
			sendPhoto(gid, replyText[6:])
		} else if strings.HasPrefix(replyText, "gif:") {
			sendGif(gid, replyText[4:])
		} else if strings.HasPrefix(replyText, "video:") {
			sendVideo(gid, replyText[6:])
		} else if strings.HasPrefix(replyText, "file:") {
			sendFile(gid, replyText[5:])
		} else if replyText != "" {
			msg = api.NewMessage(gid, replyText)
			msg.DisableWebPagePreview = true
			msg.ReplyToMessageID = upmsg.MessageID
			sendMessage(msg)
		}

		//新用户通过用户名检查是否是清真
		if upmsg.NewChatMembers != nil {
			for _, auser := range *(upmsg.NewChatMembers) {
				if checkQingzhen(auser.UserName) ||
					checkQingzhen(auser.FirstName) ||
					checkQingzhen(auser.LastName) {
					banMember(gid, uid, -1)
				}
			}
		}

		//检查清真并剔除
		if checkQingzhen(upmsg.Text) {
			_, _ = bot.DeleteMessage(api.NewDeleteMessage(gid, upmsg.MessageID))
			banMember(gid, uid, -1)
		}
	}
}

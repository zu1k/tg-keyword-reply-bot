package main

import (
	"strconv"
	"time"
	"unicode"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

/**
 * 检查是否是群组的管理员
 */
func checkAdmin(gid int64, user api.User) bool {
	admins, _ := bot.GetChatAdministrators(api.ChatConfig{ChatID: gid, SuperGroupUsername: ""})
	uid := user.ID
	if superUserId > 0 && uid == superUserId {
		return true
	}
	for _, user := range admins {
		if uid == user.User.ID {
			return true
		}
	}
	chengfa(gid, user)
	return false
}

/**
 * 检查是不是新加的群或者新开的人
 */
func checkInGroup(id int64) bool {
	for _, gid := range groups {
		if gid == id {
			return true
		}
	}
	return false
}

func chengfa(gid int64, user api.User) {
	botme, _ := bot.GetChatMember(api.ChatConfigWithUser{ChatID: gid, UserID: bot.Self.ID})
	msg := api.NewMessage(gid, "")
	if botme.CanRestrictMembers {
		banMember(gid, user.ID, 60)
		msg.Text = "[" + user.String() + "](tg://user?id=" + strconv.Itoa(user.ID) + ")乱玩管理员命令,禁言一分钟"
		msg.ParseMode = "Markdown"
	} else {
		msg.Text = "[" + user.String() + "](tg://user?id=" + strconv.Itoa(user.ID) + ")不要乱玩管理员命令"
		msg.ParseMode = "Markdown"
	}
	sendMessage(msg)
}

/**
 * 禁言群员
 */
func banMember(gid int64, uid int, sec int64) {
	if sec <= 0 {
		sec = 9999999999999
	}
	chatuserconfig := api.ChatMemberConfig{ChatID: gid, UserID: uid}
	b := false
	restricconfig := api.RestrictChatMemberConfig{
		ChatMemberConfig:      chatuserconfig,
		UntilDate:             time.Now().Unix() + sec,
		CanSendMessages:       &b,
		CanSendMediaMessages:  &b,
		CanSendOtherMessages:  &b,
		CanAddWebPagePreviews: &b}
	_, _ = bot.RestrictChatMember(restricconfig)
}

func unbanMember(gid int64, uid int) {
	chatuserconfig := api.ChatMemberConfig{ChatID: gid, UserID: uid}
	b := true
	restricconfig := api.RestrictChatMemberConfig{
		ChatMemberConfig:      chatuserconfig,
		UntilDate:             9999999999999,
		CanSendMessages:       &b,
		CanSendMediaMessages:  &b,
		CanSendOtherMessages:  &b,
		CanAddWebPagePreviews: &b}
	_, _ = bot.RestrictChatMember(restricconfig)
}

/**
 * 踢出群员
 */
func kickMember(gid int64, uid int) {
	cmconf := api.ChatMemberConfig{ChatID: gid, UserID: uid}
	_, _ = bot.KickChatMember(api.KickChatMemberConfig{ChatMemberConfig: cmconf, UntilDate: 99999999999})
}

func unkickMember(gid int64, uid int) {
	_, _ = bot.UnbanChatMember(api.ChatMemberConfig{ChatID: gid, UserID: uid})
}

/**
 * 返回群组的所有管理员, 用来进行一次性@
 */
func getAdmins(gid int64) string {
	admins, _ := bot.GetChatAdministrators(api.ChatConfig{ChatID: gid})
	list := ""
	for _, admin := range admins {
		user := admin.User
		if user.IsBot {
			continue
		}
		list += "[" + user.String() + "](tg://user?id=" + strconv.Itoa(admin.User.ID) + ")\r\n"
	}
	return list
}

/**
 * 检查文字中是否包含阿拉伯文
 */
func checkQingzhen(text string) bool {
	for _, c := range text {
		if unicode.Is(unicode.Scripts["Arabic"], c) {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
	"time"
)

/**
 * 发送文字消息
 */
func sendMessage(msg tgbotapi.MessageConfig) tgbotapi.Message {
	if msg.Text == "" {
		msg.Text = "出现错误,请联系 @veezer"
	}
	mmsg, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	go deleteMessage(msg.ChatID, mmsg.MessageID)
	return mmsg
}

/**
 * 发送图片消息, 需要是已经存在的图片链接
 */
func sendPhoto(chatid int64, photoid string) tgbotapi.Message {
	file := tgbotapi.NewPhotoShare(chatid, photoid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}
	go deleteMessage(chatid, mmsg.MessageID)
	return mmsg
}

/**
 * 发送动图, 需要是已经存在的链接
 */
func sendGif(chatid int64, gifid string) tgbotapi.Message {
	file := tgbotapi.NewAnimationShare(chatid, gifid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}
	go deleteMessage(chatid, mmsg.MessageID)
	return mmsg
}

/**
 * 发送视频, 需要是已经存在的视频连接
 */
func sendVideo(chatid int64, videoid string) tgbotapi.Message {
	file := tgbotapi.NewVideoShare(chatid, videoid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}
	go deleteMessage(chatid, mmsg.MessageID)
	return mmsg
}

/**
 * 发送文件, 必须是已经存在的文件链接
 */
func sendFile(chatid int64, fileid string) tgbotapi.Message {
	file := tgbotapi.NewDocumentShare(chatid, fileid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}
	go deleteMessage(chatid, mmsg.MessageID)
	return mmsg
}

/**
 * 转发消息, 必须是已经存在的消息链接
 */
//TODO
func forwardMessage(chatid int64, msgurl string) tgbotapi.Message {
	//https://t.me/c/1472018167/55691
	//https://t.me/cy6ersec/7
	var forward tgbotapi.ForwardConfig
	infos := strings.Split(msgurl, "/")
	var msgfromid int64
	var msgid int
	if strings.Contains(msgurl, "/c/") {
		msgid, _ = strconv.Atoi(infos[5])
		msgfromid, _ = strconv.ParseInt(infos[4], 10, 64)
		forward = tgbotapi.NewForward(chatid, msgfromid, msgid)
	} else {
		msgid, _ = strconv.Atoi(infos[4])
		forward = tgbotapi.NewForward2(chatid, infos[3], msgid)
	}

	fmt.Println(forward)
	mmsg, err := bot.Send(forward)
	if err != nil {
		log.Println(err)
	}
	return mmsg
}

func deleteMessage(gid int64, mid int) {
	time.Sleep(time.Second * 240)
	_, _ = bot.DeleteMessage(tgbotapi.NewDeleteMessage(gid, mid))
}

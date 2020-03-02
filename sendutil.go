package main

import (
	"log"
	"time"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

/**
 * 发送文字消息
 */
func sendMessage(msg api.MessageConfig) api.Message {
	if msg.Text == "" {
		return api.Message{}
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
func sendPhoto(chatid int64, photoid string) api.Message {
	file := api.NewPhotoShare(chatid, photoid)
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
func sendGif(chatid int64, gifid string) api.Message {
	file := api.NewAnimationShare(chatid, gifid)
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
func sendVideo(chatid int64, videoid string) api.Message {
	file := api.NewVideoShare(chatid, videoid)
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
func sendFile(chatid int64, fileid string) api.Message {
	file := api.NewDocumentShare(chatid, fileid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}
	go deleteMessage(chatid, mmsg.MessageID)
	return mmsg
}

func deleteMessage(gid int64, mid int) {
	time.Sleep(time.Second * 240)
	_, _ = bot.DeleteMessage(api.NewDeleteMessage(gid, mid))
}

package plugins

import (
	"github.com/3343780376/leafBot" //nolint:gci
	"github.com/3343780376/leafBot/message"
	"strconv"
	"time" //nolint:gci
)

/*
	当获取到闪照信息之后，
	会向提供的qq号进行转发该闪照
*/
func UseFlashImage(userID int) {
	leafBot.OnMessage("").SetPluginName("闪照拦截").AddRule(FlashMessageRule).AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		if userID == 0 {
			userID = leafBot.DefaultConfig.Admin
		}
		var mess message.MessageSegment
		if event.MessageType == "group" {
			mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自群\n" + strconv.Itoa(event.GroupId) + "\n的用户昵称为 " +
				event.Sender.NickName + " \nID为\n " + strconv.Itoa(event.UserId) + " \n的用户所发闪照")
		} else {
			mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自私聊,用户昵称为 " +
				event.Sender.NickName + ",ID为\n " + strconv.Itoa(event.UserId) + "\n所发闪照")
		}
		bot.SendPrivateMsg(userID, []message.MessageSegment{mess, event.Message[0].Delete("type")})
	})
}

func UseFlashImageToGroup(groupID int) {
	leafBot.OnMessage("").AddRule(FlashMessageRule).SetPluginName("闪照拦截").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		var mess message.MessageSegment
		if event.MessageType == "group" {
			mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自群\n" + strconv.Itoa(event.GroupId) + "\n的用户昵称为 " +
				event.Sender.NickName + " \nID为\n " + strconv.Itoa(event.UserId) + " \n的用户所发闪照")
		} else {
			mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自私聊,用户昵称为 " +
				event.Sender.NickName + ",ID为\n " + strconv.Itoa(event.UserId) + "\n所发闪照")
		}
		bot.SendGroupMsg(groupID, []message.MessageSegment{mess, event.Message[0].Delete("type")})
	})

}

func FlashMessageRule(event leafBot.Event, bot *leafBot.Bot) bool {
	for _, msg := range event.GetMsg() {
		if msg.Type == "image" && msg.Data["type"] == "flash" {
			return true
		}
	}
	return false
}

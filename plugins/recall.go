package plugins

import (
	"strconv"
	"time" //nolint:gci

	"github.com/3343780376/leafBot" //nolint:gci
	"github.com/3343780376/leafBot/message"
)

/*
	当获取到撤回信息之后，
	会向提供的qq号进行转发该撤回信息
*/

func UseRecall(userID int) {
	leafBot.OnNotice("group_recall").SetPluginName("群组内撤回发给自己").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		var mess message.MessageSegment
		mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自群\n" + strconv.Itoa(event.GroupId) +
			" \nID为\n " + strconv.Itoa(event.UserId) + " \n的用户所撤回信息\n")

		bot.SendPrivateMsg(userID, append([]message.MessageSegment{mess}, bot.GetMsg(event.MessageID).Message...))

	})

	leafBot.OnNotice("friend_recall").SetPluginName("朋友撤回发给自己").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		var mess message.MessageSegment
		mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自私聊,ID为\n " + strconv.Itoa(event.UserId) + "\n所撤回信息\n")

		bot.SendPrivateMsg(userID, append([]message.MessageSegment{mess}, bot.GetMsg(event.MessageID).Message...))

	})
}

func UseRecallToGroup(groupID int) {

	leafBot.OnNotice("group_recall").SetPluginName("群组内撤回发给群组").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {

		if event.GroupId == 526408479 || event.UserId == 3269784826 || event.GroupId == 1048452984 || event.UserId == 2457444952 {
			return
		}

		var mess message.MessageSegment
		mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自群\n" + strconv.Itoa(event.GroupId) +
			" \nID为\n " + strconv.Itoa(event.UserId) + " \n的用户所撤回信息\n")

		bot.SendGroupMsg(groupID, append([]message.MessageSegment{mess}, bot.GetMsg(event.MessageID).Message...))

	})

	leafBot.OnNotice("friend_recall").SetPluginName("朋友撤回发给群组").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		var mess message.MessageSegment
		mess = message.Text(time.Now().Format("2006-01-02 15:04:05") + "\n来自私聊 ,ID为\n " + strconv.Itoa(event.UserId) + "\n所撤回信息\n")

		bot.SendGroupMsg(groupID, append([]message.MessageSegment{mess}, bot.GetMsg(event.MessageID).Message...))
	})
}

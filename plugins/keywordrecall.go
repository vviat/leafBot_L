package plugins

import (
	"strconv"
	"strings"
	"time" //nolint:gci

	"github.com/3343780376/leafBot" //nolint:gci
)

/*
	触发关键词 自动撤回
*/
func KeyWordRecall(userID int) {
	leafBot.OnMessage("").SetPluginName("关键词撤回").AddRule(KeyWordRecallRule).AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		if event.GroupId == 797211774 || event.GroupId == 623908643 || event.GroupId == 693646618 || event.GroupId == 232226765 {
			time.Sleep(2 * time.Second)
			bot.SetGroupBan(event.GroupId, event.UserId, 1*60)
			bot.DeleteMsg(event.MessageID)
			mess := "该消息包括敏感关键词，请注意言行\n"
			allmess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(20s后撤回)\n" + mess
			MessageID := bot.SendGroupMsg(event.GroupId, allmess)
			time.Sleep(20 * time.Second)
			bot.DeleteMsg(MessageID)
		}
	})
}

func KeyWordRecallRule(event leafBot.Event, bot *leafBot.Bot) bool {
	words := []string{"新生墙", "资料墙", "校园墙", "新生必备墙", "墙墙", "指南墙", "好评"}
	for _, word := range words {
		eventmess := event.Message.CQString()
		if strings.Contains(eventmess, word) {
			return true
		}
	}
	return false
}

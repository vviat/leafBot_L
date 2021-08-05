package plugins

import (
	"strconv"
	"time"

	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
)

func UseEchoHandle() {

	leafBot.
		OnCommand("/vvecho").
		SetPluginName("vvecho").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(args[0]))
			})

	leafBot.
		OnCommand("/老乡群").
		AddAllies("/lxq").
		AddAllies("lxq").
		AddAllies("老乡群").
		SetPluginName("老乡群").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				mess := "2021陕西老乡会：872484216\n2021四川老乡会：160284029\n2021山西老乡会：630427365\n2021云南老乡会：753984336\n2021海南老乡会：309107584\n2021湖北老乡会：1136509440\n2021甘肃老乡会：637987762\n2021河南老乡会：790379210\n2021安徽老乡会：895675852\n2021河北老乡会：643987089\n2021浙江老乡会：228948266\n2021广西老乡会：367881633\n2021吉林老乡会：662129408\n2021江苏老乡会：789901430\n2021山东老乡会：436850757\n"
				mianzi := "信息为群内整理 真假自行分辨 需要补充私聊我\n"
				allmess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(40s后撤回)\n" + mess + mianzi
				MessageID := bot.SendGroupMsg(event.GroupId, allmess)
				time.Sleep(40 * time.Second)
				bot.DeleteMsg(MessageID)
			})

	leafBot.
		OnCommand("/寝室").
		AddAllies("/qs").
		SetPluginName("寝室").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				mess := "[CQ:image,file=file:////home/viat/vv_images/image/qs.png]\n"
				//mianzi := "信息为群内整理 真假自行分辨 需要补充私聊我\n"
				allmess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(40s后撤回)\n" + mess
				MessageID := bot.SendGroupMsg(event.GroupId, allmess)
				time.Sleep(40 * time.Second)
				bot.DeleteMsg(MessageID)
			})
	leafBot.
		OnCommand("/社团").
		AddAllies("/st").
		SetPluginName("社团").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				mess := "[CQ:image,file=file:////home/viat/vv_images/image/hdst.png]\n"
				//mianzi := "信息为群内整理 真假自行分辨 需要补充私聊我\n"
				allmess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(40s后撤回)\n" + mess
				MessageID := bot.SendGroupMsg(event.GroupId, allmess)
				time.Sleep(40 * time.Second)
				bot.DeleteMsg(MessageID)
			})
	leafBot.
		OnCommand("/问题一").
		AddAllies("/Q1").
		SetPluginName("问题一").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				mess := "[CQ:image,file=file:////home/viat/vv_images/image/Q1.jpeg]\n"
				//mianzi := "信息为群内整理 真假自行分辨 需要补充私聊我\n"
				allmess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(40s后撤回)\n" + mess
				MessageID := bot.SendGroupMsg(event.GroupId, allmess)
				time.Sleep(40 * time.Second)
				bot.DeleteMsg(MessageID)
			})

	leafBot.
		OnCommand("/问题二").
		AddAllies("/Q2").
		SetPluginName("问题二").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				mess := "[CQ:image,file=file:////home/viat/vv_images/image/Q2.jpeg]\n"
				//mianzi := "信息为群内整理 真假自行分辨 需要补充私聊我\n"
				allmess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(40s后撤回)\n" + mess
				MessageID := bot.SendGroupMsg(event.GroupId, allmess)
				time.Sleep(40 * time.Second)
				bot.DeleteMsg(MessageID)
			})

	leafBot.
		OnCommand("/lq帮助").
		AddAllies("/lqhelp").
		SetPluginName("帮助").
		SetWeight(1).
		SetBlock(false).
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				mess := "支持触发如下\n`/老乡群`， `/lxq`\n`/寝室`， `/qs`\n`/问题一`， `/Q1`\n`/问题二`， `/Q2`\n"
				//mianzi := "信息为群内整理 真假自行分辨 需要补充私聊我\n"
				allmess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(30s后撤回)\n" + mess
				MessageID := bot.SendGroupMsg(event.GroupId, allmess)
				time.Sleep(30 * time.Second)
				bot.DeleteMsg(MessageID)
			})

	/*
		leafBot.
			OnCommand("/精华消息").
			SetPluginName("精华消息").
			SetWeight(1).
			SetBlock(false).
			AddHandle(
				func(event leafBot.Event, bot *leafBot.Bot, args []string) {
					var AllGroupEssenceMsg message.Message
					GroupEssenceMsgList := bot.GetEssenceMsgList(event.GroupId)
					fmt.Printf("33333333333333\n")
					fmt.Printf("%v\n", GroupEssenceMsgList)
					fmt.Printf("44444444444444\n")
					for _, v := range GroupEssenceMsgList {
						GroupEssenceMsg := bot.GetMsg(int32(v.Message_id)).Message
						AllGroupEssenceMsg = append(AllGroupEssenceMsg, GroupEssenceMsg...)
						fmt.Printf("11111111111111\n")
						fmt.Printf("%s\n", AllGroupEssenceMsg)
						fmt.Printf("22222222222222\n")
					}
					mess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + "(60s后撤回)\n"
					allmess := message.ParseMessageFromString(mess)
					allmess = append(allmess, AllGroupEssenceMsg...)
					MessageID := bot.SendGroupMsg(event.GroupId, allmess)
					time.Sleep(60 * time.Second)
					bot.DeleteMsg(MessageID)
				})
	*/
}

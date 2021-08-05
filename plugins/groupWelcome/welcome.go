// Package groupWelcome /**
package groupWelcome

import (
	"strconv"
	"time"

	"github.com/3343780376/leafBot"
)

// WelcomeInit
/**
 * @Description:
 */
func WelcomeInit() {
	leafBot.OnNotice(leafBot.NoticeTypeApi.GroupIncrease).SetWeight(10).SetPluginName("入群欢迎").AddHandle(func(event leafBot.Event, bot *leafBot.Bot) {
		for _, s := range leafBot.DefaultConfig.Plugins.Welcome {
			if s.GroupId == event.GroupId {
				var member_sex string = "(无法通过QQ搜索该成员 个人信息无法获取)"
				//memberinfo := bot.GetGroupMemberInfo(event.GroupId, event.UserId, false)
				strangerinfo := bot.GetStrangerInfo(event.UserId, false)

				switch strangerinfo.Sex {
				case "male":
					member_sex = "欢迎学弟"
				case "female":
					member_sex = "欢迎学妹"
				case "unknown":
					member_sex = "欢迎新同学"
				default:
					member_sex = "(无法通过QQ搜索该成员 信息无法获取)"
				}

				mess := "[CQ:at,qq=" + strconv.Itoa(event.UserId) + "]" + "\n" + member_sex + "(90s后撤回)\n"
				mess += s.Message
				MessageID := bot.SendGroupMsg(event.GroupId, mess)
				time.Sleep(90 * time.Second)
				bot.DeleteMsg(MessageID)
			}
		}
	})
}

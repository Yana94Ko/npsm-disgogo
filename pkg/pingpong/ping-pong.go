package pingpong

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"npsm_gogo/config"
)

func HandlePingPong(s *discordgo.Session) {
	s.AddHandler(PingPong)
}

func PingPong(s *discordgo.Session, m *discordgo.MessageCreate) {
	e := config.LoadEnvVars()
	var chatKeyValueMap = map[string]string{
		"노프": "스모!",
		"찬웅": fmt.Sprintf("ㅋㅋ 제 블로그 보실? %s", e.ChanWoongBlog),
		"야나": fmt.Sprintf("ㅋㅋ 제 블로그 보실? %s", e.YanaBlog),
	}
	if m.Author.ID == s.State.User.ID {
		return
	}
	if reply, exists := chatKeyValueMap[m.Content]; exists {
		s.ChannelMessageSend(m.ChannelID, reply)
	}
}

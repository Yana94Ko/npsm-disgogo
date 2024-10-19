package scheduler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	"time"
)

func CreateScheduler(discord *discordgo.Session, channelID string, schedule string) {
	c := cron.New(cron.WithSeconds())

	location, error := time.LoadLocation("Asia/Seoul")
	if error != nil {
		log.Logger.Err(error).Msg("Error loading location")
	}

	_, err := c.AddFunc(schedule, func() {
		now := time.Now().In(location)

		msg := "@here :timer: :timer: 스터디 시간이 다가왔습니다! 준비해주세요. \n :coffee: :coffee: :coffee:  온라인인데도 늦으면 커피사겠지 뭐 ㅋ"
		log.Logger.Debug().Msgf("kr_time : %v, msg : %s", now, msg)

		_, err := discord.ChannelMessageSend(channelID, msg)
		if err != nil {
			log.Logger.Err(err).Msgf("Error sending scheduled message: %v", err)
		}
	})
	if err != nil {
		log.Logger.Err(err).Msgf("Error adding cron job: %v", err)
	}

	c.Start()

	log.Logger.Info().Msgf("Scheduler started with schedule: %s", schedule)
}

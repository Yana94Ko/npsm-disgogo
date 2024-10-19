package main

import (
	"crypto/tls"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"npsm_gogo/config"
	"npsm_gogo/pkg/pingpong"
	"npsm_gogo/pkg/scheduler"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	e := config.LoadEnvVars()

	log.Logger.Info().Msg("Started Discord Bot [[ NPSM_GOGO ]]")

	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	dialer := &net.Dialer{}

	discord, err := discordgo.New("Bot " + e.Token)
	if err != nil {
		log.Logger.Err(err).Msgf("Error creating Discord session, error : %v", err)
		return
	}
	defer discord.Close()

	discord.Dialer = &websocket.Dialer{
		NetDial:          dialer.Dial,
		TLSClientConfig:  tlsConfig,
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	pingpong.HandlePingPong(discord)

	scheduler.CreateScheduler(discord, e.DefaultChannel, "0 55 9 * * 6")

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		log.Logger.Err(err).Msgf("Error opening connection, error : %v", err)
		return
	}

	log.Logger.Info().Msg("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

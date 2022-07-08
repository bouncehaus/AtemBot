package mux

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) Ping(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	_, err := ds.ChannelMessageSend(dm.ChannelID, "<@$USER_ID>")

	time.Sleep(2 * time.Second)

	ds.ChannelMessageSend(dm.ChannelID, "<@$USER_ID>")

	time.Sleep(2 * time.Second)

	ds.ChannelMessageSend(dm.ChannelID, "<@$USER_ID>")

	time.Sleep(2 * time.Second)

	ds.ChannelMessageSend(dm.ChannelID, "<@$USER_ID>")

	time.Sleep(2 * time.Second)

	ds.ChannelMessageSend(dm.ChannelID, "<@$USER_ID>")
	if err != nil {
		log.Print("err", err)

	}

}

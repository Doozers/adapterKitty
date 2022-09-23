package main

import (
	"fmt"
	"os"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
	"github.com/bwmarrin/discordgo"
)

type discordSvc struct {
	Plugin func([]byte)
	Token  string
	dg     *discordgo.Session
}

func (opt *discordSvc) listener(client proto.Serv_AdapterClient) {
	dg, err := discordgo.New("Bot " + opt.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		os.Exit(1)
	}

	opt.dg = dg

	opt.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			fmt.Println("Message received: ", m.Content)
		}
		if err := client.Send(&proto.AdapterRequest{Payload: []byte(m.Content)}); err != nil {
			fmt.Println("Error1: ", err)
			return
		}
	})

	fmt.Println("Bot is running")
	err = opt.dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

func (opt *discordSvc) react(b []byte) {
	if opt.Plugin != nil {
		opt.Plugin(b)
		return
	}

	// default reaction
	fmt.Print("LOGS: SERV ANSWER >> ", string(b), "\n\n >> ")
	opt.dg.ChannelMessageSend("", string(b))
}

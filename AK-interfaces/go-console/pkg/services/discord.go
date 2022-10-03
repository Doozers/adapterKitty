package services

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
	"github.com/bwmarrin/discordgo"
)

type DiscordSvc struct {
	FormatPlug func([]byte) ([]byte, error)
	ReactPlug  func([]byte) (string, error)

	Token string
	dg    *discordgo.Session
	Chan  string

	Type GrpcType
}

func (svc *DiscordSvc) UniListener(ctx context.Context, client proto.AdapterKitServiceClient) {
	dg, err := discordgo.New("Bot " + svc.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	svc.dg = dg
	err = svc.dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	defer svc.dg.Close()

	svc.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		fmt.Println("Message received: ", m.Content)
		b, err := svc.Format([]byte(m.Content))
		if err != nil {
			return
		}
		resp, err := client.UniDirectionalAdapter(ctx, &proto.AdapterRequest{Payload: b})
		if err != nil {
			fmt.Println("error sending request:", err)
			return
		}

		r, err := svc.React(resp.GetPayload())
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = s.ChannelMessageSendReply(m.ChannelID, r, m.Reference())
		if err != nil {
			fmt.Println("error replying", err)
			return
		}
	})

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func (svc *DiscordSvc) SsListener(ctx context.Context, client proto.AdapterKitServiceClient) {
	dg, err := discordgo.New("Bot " + svc.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	svc.dg = dg
	err = svc.dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	establishedStream := make(chan int, 1)
	var stream proto.AdapterKitService_ServerStreamingAdapterClient

	svc.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		fmt.Println("Message received: ", m.Content)
		b, err := svc.Format([]byte(m.Content))
		svc.Chan = m.ChannelID
		if err != nil {
			return
		}

		stream, err = client.ServerStreamingAdapter(ctx, &proto.AdapterRequest{Payload: b})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		establishedStream <- 1
	})

	fmt.Println("Discord bot is now running.  Press CTRL-C to exit.")
	<-establishedStream
	svc.dg.Close()
	fmt.Println("Stream established")

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error23: ", err)
				return
			}
			r, err := svc.React(resp.Payload)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = dg.ChannelMessageSend(svc.Chan, r)
			if err != nil {
				fmt.Println("error sending message on discord", err)
				return
			}
		}
	}()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func (svc *DiscordSvc) BiListener(client proto.AdapterKitService_BiDirectionalAdapterClient) {

}

func (svc *DiscordSvc) GetType() GrpcType {
	return svc.Type
}

func (svc *DiscordSvc) Format(msg []byte) ([]byte, error) {
	if svc.FormatPlug != nil {
		res, err := svc.FormatPlug(msg)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return msg, nil
}

func (svc *DiscordSvc) React(b []byte) (string, error) {
	if svc.ReactPlug != nil {
		res, err := svc.ReactPlug(b)
		if err != nil {
			return "", err
		}
		return res, nil
	}

	return string(b), nil
}

package discord

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"syscall"

// 	"github.com/bwmarrin/discordgo"
// )

// const REACTION_SUCCEED = "✅"
// const REACTION_FAILED = "❌"

// func NewDiscord() {
// 	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dg.AddHandler(messageCreate)
// 	err = dg.Open()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	sc := make(chan os.Signal, 1)
// 	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
// 	<-sc

// 	dg.Close()
// }

// func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	prefix := os.Getenv("DISCORD_PREFIX")

// 	if m.Author.ID == s.State.User.ID || m.Author.Bot || !strings.HasPrefix(m.Content, prefix) {
// 		return
// 	}

// 	cmd := strings.Replace(m.Content, prefix, "", 1)
// 	args := strings.Split(cmd, " ")

// 	if strings.HasPrefix(cmd, "session") {
// 		switch args[1] {
// 		case "create":
// 			url := auth.AuthURL(os.Getenv("SPOTIFY_OAUTH_STATE"))
// 			s.ChannelMessageSendReply(m.ChannelID,
// 				fmt.Sprintf("極度認証(しなさい) %s", url),
// 				m.Reference())
// 		case "revoke":
// 		}
// 	}

// 	// Start Spotify
// 	if spotifyClient == nil {
// 		s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_FAILED)
// 		s.ChannelMessageSendReply(m.ChannelID, "**😲NOT AUTHENTICATED**", m.Reference())
// 		return
// 	}

// 	if strings.HasPrefix(cmd, "play") {
// 		if len(args) == 1 {
// 			err := spotifyClient.Play()
// 			if err != nil {
// 				log.Println(err)
// 				s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_FAILED)
// 			} else {
// 				s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_SUCCEED)
// 				s.ChannelMessageSendReply(m.ChannelID, "**🎵START PLAYING**", m.Reference())
// 			}
// 		} else {
// 			res, err := spotifyClient.Search(args[1], spotify.SearchTypeTrack)
// 			if err != nil {
// 				log.Println(err)
// 				s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_FAILED)
// 			}
// 			log.Println(res.Tracks.Tracks[0])
// 			opt := &spotify.PlayOptions{
// 				DeviceID: nil,
// 			}

// 			err = spotifyClient.QueueSongOpt(res.Tracks.Tracks[0].ID, opt)
// 			if err != nil {
// 				log.Println(err)
// 				s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_FAILED)
// 			} else {
// 				s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_SUCCEED)
// 				s.ChannelMessageSendReply(m.ChannelID,
// 					fmt.Sprintf("**💿UPCOMING TRACK:** %s(%s)",
// 						res.Tracks.Tracks[0].Name,
// 						res.Tracks.Tracks[0].Album.Name),
// 					m.Reference())
// 			}
// 		}
// 	}

// 	if strings.HasPrefix(cmd, "skip") {
// 		err := spotifyClient.Next()
// 		if err != nil {
// 			log.Println(err)
// 			s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_FAILED)
// 		} else {
// 			s.MessageReactionAdd(m.ChannelID, m.ID, REACTION_SUCCEED)
// 			s.ChannelMessageSendReply(m.ChannelID, "**⏭️SKIPPED**", m.Reference())
// 		}
// 	}
// }

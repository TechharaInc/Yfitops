// クッキンアイドル アイ!マイ!まいん! - Wikipedia
// https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%83%E3%82%AD%E3%83%B3%E3%82%A2%E3%82%A4%E3%83%89%E3%83%AB_%E3%82%A2%E3%82%A4!%E3%83%9E%E3%82%A4!%E3%81%BE%E3%81%84%E3%82%93!

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
)

var auth spotify.Authenticator

var (
	ch = make(chan *spotify.Client)
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("AAGH .env COULD NOT LOADED!")
	}

	auth = spotify.NewAuthenticator(os.Getenv("REDIRECT_URL"), spotify.ScopeUserReadPrivate)
	auth.SetAuthInfo(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"))

	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("WTF", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	log.Println("Server start listening at :8080")

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	dg.Close()
}

func awaitCompleteAuth() {
	client := <-ch

	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WE ARE LOGGED IN AS:", user.DisplayName)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(os.Getenv("SPOTIFY_OAUTH_STATE"), r)
	if err != nil {
		log.Printf("%s\n", err.Error())
		http.Error(w, "COULDN'T GET TOKEN!", http.StatusNotFound)
		return
	}
	if st := r.FormValue("state"); st != os.Getenv("SPOTIFY_OAUTH_STATE") {
		http.NotFound(w, r)
		return
	}
	client := auth.NewClient(token)
	fmt.Fprintf(w, "このページは閉じていいゾ")
	ch <- &client
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	cmd := strings.Replace(m.Content, os.Getenv("DISCORD_PREFIX"), "", 1)
	args := strings.Split(cmd, " ")

	if strings.HasPrefix(cmd, "session") {
		switch args[1] {
		case "create":
			url := auth.AuthURL(os.Getenv("SPOTIFY_OAUTH_STATE"))
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("極度認証(しなさい) %s", url))

		}
	}
}

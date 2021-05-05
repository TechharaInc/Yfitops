package client

import (
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type SpotifyClient struct {
	auth spotify.Authenticator
}

func NewSpotifyClient() (*SpotifyClient, error) {
	auth := spotify.NewAuthenticator(
		os.Getenv("REDIRECT_URL"),
		spotify.ScopeUserReadPrivate,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserModifyPlaybackState)
	auth.SetAuthInfo(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"))

	return &SpotifyClient{
		auth,
	}, nil
}

func (c *SpotifyClient) GetAuthURL(state string) string {
	return c.auth.AuthURL(state)
}

func (c *SpotifyClient) Exchange(code string) (*oauth2.Token, error) {
	token, err := c.auth.Exchange(code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (c *SpotifyClient) Refresh(token *oauth2.Token) (*oauth2.Token, error) {
	cli := c.auth.NewClient(token)
	newToken, err := cli.Token()
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

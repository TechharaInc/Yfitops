package client

import (
	"context"
	"errors"

	"github.com/TechharaInc/Yfitops/service"
	"github.com/zmb3/spotify"
)

func (sc *SpotifyClient) Search(ctx context.Context, query string) (*spotify.SearchResult, error) {
	token, ok := service.GetTokenFromContext(ctx)
	if !ok {
		return nil, errors.New("TOKEN NOT FOUND:astonished:")
	}
	cli := sc.auth.NewClient(token)
	res, err := cli.Search(query, spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (sc *SpotifyClient) Play(ctx context.Context) error {
	token, ok := service.GetTokenFromContext(ctx)
	if !ok {
		return nil
	}
	cli := sc.auth.NewClient(token)
	return cli.Play()
}

func (sc *SpotifyClient) QueueSongOpt(ctx context.Context, trackID spotify.ID, opt *spotify.PlayOptions) error {
	token, ok := service.GetTokenFromContext(ctx)
	if !ok {
		return nil
	}
	cli := sc.auth.NewClient(token)
	return cli.QueueSongOpt(trackID, opt)
}

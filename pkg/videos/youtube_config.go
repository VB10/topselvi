package videos

import (
	"context"
	"errors"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const youtubeEP = "YOUTUBE_ENDPOINT"

// YoutubeConfig take video config.
func YoutubeVideoDetail(id string) (*youtube.Video, error) {
	apiKey, _ := os.LookupEnv("YOUTUBE_API_KEY")
	ctx := context.Background()
	youtubeService, error := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if error != nil {
		return nil, error
	}

	result, error := youtubeService.Videos.List("snippet,contentDetails,statistics").Id(id).Do()
	if len(result.Items) <= 0 {
		return nil, errors.New("Error.Channel id not found")
	}

	return result.Items[0], nil
}

//  YoutubeUserDetail get  Youtube user detail.
func YoutubeUserDetail(id string) (*User, error) {
	youtubeVideo, error := YoutubeVideoDetail(id)
	if error != nil {
		return nil, error
	}
	apiKey, _ := os.LookupEnv("YOUTUBE_API_KEY")
	ctx := context.Background()
	youtubeService, error := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	resultChannel, error := youtubeService.Channels.List("snippet,contentDetails,statistics").Id(youtubeVideo.Snippet.ChannelId).Do()
	if len(resultChannel.Items) <= 0 {
		print("Error.Channel id not found")
		return nil, error
	}

	var user User
	user.Username = resultChannel.Items[0].Snippet.Title
	user.ProfileImage = resultChannel.Items[0].Snippet.Thumbnails.Default.Url
	return &user, nil
}

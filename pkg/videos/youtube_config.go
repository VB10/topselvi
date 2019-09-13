package videos

import (
	"context"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const youtubeEP = "YOUTUBE_ENDPOINT"

// YoutubeConfig take video config.
func YoutubeConfig(id string) *User {
	apiKey, _ := os.LookupEnv("YOUTUBE_API_KEY")
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		println("Youtube service error" + err.Error())
		return nil
	}

	result, _ := youtubeService.Videos.List("snippet,contentDetails,statistics").Id(id).Do()
	if len(result.Items) <= 0 {
		print("Error.Channel id not found")
		return nil
	}
	var youtubeVideo = result.Items[0]

	resultChannel, _ := youtubeService.Channels.List("snippet,contentDetails,statistics").Id(youtubeVideo.Snippet.ChannelId).Do()
	if len(result.Items) <= 0 {
		print("Error.Channel id not found")
		return nil
	}


	var user User
	user.Username = youtubeVideo.Snippet.ChannelTitle
	user.ProfileImage = resultChannel.Items[0].Snippet.Thumbnails.Default.Url
	return &user
}

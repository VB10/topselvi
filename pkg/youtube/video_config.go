package youtube

import (
	"context"
	"os"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const youtubeEP = "YOUTUBE_ENDPOINT"

// YoutubeConfig take video config.
func YoutubeConfig(id string) *youtube.Video {
	apiKey, _ := os.LookupEnv("YOUTUBE_API_KEY")
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		// TODO : LOGGING SERVICE ADD
		println("Youtube service error" + err.Error())
	}
	result, _ := youtubeService.Videos.List("snippet,contentDetails,statistics").Id(id).Do()
	if len(result.Items) <= 0 {
		print("Error.Channel id not found")
		return nil
	}

	print(result.Items[0].Snippet.Title)
	return result.Items[0]
}

//VideoConfig is Youtube video detail model.
type VideoConfig struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string   `json:"channelTitle"`
			Tags                 []string `json:"tags"`
			CategoryID           string   `json:"categoryId"`
			LiveBroadcastContent string   `json:"liveBroadcastContent"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
		} `json:"snippet"`
		ContentDetails struct {
			Duration        string `json:"duration"`
			Dimension       string `json:"dimension"`
			Definition      string `json:"definition"`
			Caption         string `json:"caption"`
			LicensedContent bool   `json:"licensedContent"`
			Projection      string `json:"projection"`
		} `json:"contentDetails"`
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			DislikeCount  string `json:"dislikeCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
}

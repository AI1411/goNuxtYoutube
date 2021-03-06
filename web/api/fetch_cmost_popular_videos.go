package api

import (
	"context"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		//keyを設定
		key := os.Getenv("API_KEY")

		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(key))
		if err != nil {
			logrus.Fatalf("Error creating new Youtube service: %v", err)
		}
		call := yts.Videos.List("id,snippet").Chart("mostPopular").MaxResults(3)

		//続きを表示する用の処理
		pageToken := c.QueryParam("pageToken")
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}
		//APIを実行
		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		return c.JSON(fasthttp.StatusOK, res)
	}
}

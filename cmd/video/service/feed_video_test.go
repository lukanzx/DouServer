package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/lukanzx/DouServer/cmd/video/dal"
	"github.com/lukanzx/DouServer/cmd/video/rpc"
	"github.com/lukanzx/DouServer/cmd/video/service"
	"github.com/lukanzx/DouServer/config"
	"github.com/lukanzx/DouServer/kitex_gen/video"
	"github.com/lukanzx/DouServer/pkg/utils"
)

func TestFeedVideo(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	token, err := utils.CreateToken(10000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	videoService := service.NewVideoService(context.Background())
	testTime := new(int64)
	*testTime = 1693101739
	// 第一次从数据库读取，并写入redis
	_, _, _, _, _, err = videoService.FeedVideo(&video.FeedRequest{
		LatestTime: testTime,
		Token:      &token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	// 等待写入redis
	time.Sleep(time.Second)
	// 第二次直接从redis读取
	_, _, _, _, _, err = videoService.FeedVideo(&video.FeedRequest{
		LatestTime: testTime,
		Token:      &token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

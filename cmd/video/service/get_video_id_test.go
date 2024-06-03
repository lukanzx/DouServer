package service_test

import (
	"context"
	"testing"

	"github.com/lukanzx/DouServer/cmd/video/dal"
	"github.com/lukanzx/DouServer/cmd/video/rpc"
	"github.com/lukanzx/DouServer/cmd/video/service"
	"github.com/lukanzx/DouServer/config"
	"github.com/lukanzx/DouServer/kitex_gen/video"
	"github.com/lukanzx/DouServer/pkg/utils"
)

func TestGetVideoIDByUid(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	token, err := utils.CreateToken(10000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	videoService := service.NewVideoService(context.Background())
	_, err = videoService.GetVideoIDByUid(&video.GetVideoIDByUidRequset{Token: token, UserId: 10000})

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

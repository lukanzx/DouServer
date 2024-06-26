package main

import (
	"testing"

	"github.com/lukanzx/DouServer/kitex_gen/video"
)

func testGetFavoriteVideo(t *testing.T) {
	_, _, _, _, err := videoService.GetFavoriteVideoInfo(&video.GetFavoriteVideoInfoRequest{
		VideoId: videoId,
		Token:   token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

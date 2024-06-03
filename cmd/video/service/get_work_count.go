package service

import (
	"github.com/lukanzx/DouServer/cmd/video/dal/db"
	"github.com/lukanzx/DouServer/kitex_gen/video"
)

func (s *VideoService) GetWorkCount(req *video.GetWorkCountRequest) (workCount int64, err error) {
	return db.GetWorkCountByUid(s.ctx, req.UserId)
}

package service

import (
	"errors"

	"github.com/lukanzx/DouServer/cmd/interaction/dal/cache"
	"github.com/lukanzx/DouServer/cmd/interaction/dal/db"
	"github.com/lukanzx/DouServer/kitex_gen/interaction"
	"github.com/lukanzx/DouServer/pkg/errno"
	"gorm.io/gorm"
)

func (s *InteractionService) Dislike(req *interaction.FavoriteActionRequest, userID int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userID)
	if err != nil {
		return err
	}
	if !exist {
		err := db.IsFavorited(s.ctx, userID, req.VideoId, 1)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LikeNoExistError
		}
	}

	ok, _, err := cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		return err
	}
	if ok {
		if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId, userID); err != nil {
			return err
		}
	}

	// write into mysql
	return db.UpdateFavoriteStatus(s.ctx, userID, req.VideoId, 0)
}

package dal

import (
	"github.com/lukanzx/DouServer/cmd/video/dal/cache"
	"github.com/lukanzx/DouServer/cmd/video/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
